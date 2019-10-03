package rwacryptocell

import (
	"errors"
	fmt "fmt"
	"sort"
)

func ValidateContents(who *Who, next *CellContents) []error {
	var errs []error

	if who == nil {
		return nil
	}
	errs2 := ValidateWho(who, next.Who, next.WhoSigs)
	errs = append(errs, errs2...)

	errs2 = ValidateWhat(next.Who, next.What, next.WhatAuthor, next.WhatSig)
	errs = append(errs, errs2...)

	return errs
}

func ValidateWhat(who *Who, what *What, signerIndex int32, sig *Sig) []error {
	var errs []error

	if !int32Contains(who.Write, signerIndex) {
		err := fmt.Errorf("author does not have write permission")
		errs = append(errs, err)
	}

	author := who.Entities[int(signerIndex)]

	if what == nil {
		what = &What{}
	}
	whatBytes, err := what.XXX_Marshal(nil, true)
	if err != nil {
		panic(err)
	}

	if err := VerifySig(whatBytes, author.SigningKey, sig); err != nil {
		errs = append(errs, err)
	}

	return nil
}

func ValidateWho(prev, next *Who, sigs map[int32]*Sig) []error {
	var errs []error
	if next == nil {
		err := errors.New("nil who")
		errs = append(errs, err)
		return errs
	}

	admins := prev.Admin
	sort.Slice(admins, func(i, j int) bool {
		return admins[i] < admins[j]
	})

	whoBytes, err := next.XXX_Marshal(nil, true)
	if err != nil {
		panic(err)
	}

	// must find an admin who signed off on it
	for i, sig := range sigs {
		if int32Contains(admins, i) {
			signer := prev.Entities[i].SigningKey
			if err := VerifySig(whoBytes, signer, sig); err != nil {
				errs = append(errs, err)
			}
			return errs
		}
	}
	noSigErr := fmt.Errorf("no valid admin sig found")
	errs = append(errs, noSigErr)
	return errs
}

func int32Contains(s []int32, x int32) bool {
	n := sort.Search(len(s), func(i int) bool {
		return x >= s[i]
	})
	if n >= len(s) {
		return false
	}
	return s[n] == x
}
