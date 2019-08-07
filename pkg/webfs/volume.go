package webfs

import (
	"context"
	"errors"

	"github.com/brendoncarroll/webfs/pkg/cells"
	"github.com/brendoncarroll/webfs/pkg/webfs/models"
	"github.com/brendoncarroll/webfs/pkg/webref"
)

const VolumeCodec = webref.CodecJSON

var ErrCASFailed = errors.New("CAS failed. Should Retry")

type Cell = cells.Cell

type VolumeMutator func(models.Volume) (*models.Volume, error)
type ObjectMutator func(*models.Object) (*models.Object, error)

type Volume struct {
	cell Cell
	opts *Options

	baseObject
}

func (v *Volume) Find(ctx context.Context, p Path, objs []Object) ([]Object, error) {
	if len(p) == 0 {
		objs = append(objs, v)
	}

	o, err := v.getObject(ctx)
	if err != nil {
		return nil, err
	}
	if o != nil {
		objs, err = o.Find(ctx, p, objs)
		if err != nil {
			return nil, err
		}
	}
	return objs, nil
}

func (v *Volume) Lookup(ctx context.Context, p Path) (Object, error) {
	o, err := v.getObject(ctx)
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, nil
	}
	return o.Lookup(ctx, p)
}

func (v *Volume) Walk(ctx context.Context, f func(Object) bool) (bool, error) {
	cont := f(v)
	if !cont {
		return false, nil
	}

	cc, err := v.Get(ctx)
	if err != nil {
		return false, err
	}
	mo := models.Object{}
	s := v.getStore()
	if err := webref.Load(ctx, s, *cc.ObjectRef, &mo); err != nil {
		return false, err
	}
	o, err := wrapObject(v, "", &mo)
	if err != nil {
		return false, err
	}

	return o.Walk(ctx, f)
}

func (v *Volume) Get(ctx context.Context) (*models.Volume, error) {
	data, err := v.cell.Get(ctx)
	if err != nil {
		return nil, err
	}
	if len(data) < 1 {
		return &models.Volume{
			Options: DefaultOptions(),
		}, nil
	}

	m := models.Volume{}
	if err := webref.Decode(webref.CodecJSON, data, &m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (v *Volume) getObject(ctx context.Context) (Object, error) {
	m, err := v.Get(ctx)
	if err != nil {
		return nil, err
	}
	if m.ObjectRef == nil {
		return nil, nil
	}
	mo := models.Object{}
	s := v.getStore()
	if err := webref.Load(ctx, s, *m.ObjectRef, &mo); err != nil {
		return nil, err
	}
	return wrapObject(v, "", &mo)
}

func (v *Volume) put(ctx context.Context, fn ObjectMutator) error {
	return v.apply(ctx, func(cur models.Volume) (*models.Volume, error) {
		var o *models.Object
		if cur.ObjectRef != nil {
			o = &models.Object{}
			s := v.getStore()
			if err := webref.Load(ctx, s, *cur.ObjectRef, o); err != nil {
				return nil, err
			}
		}

		o2, err := fn(o)
		if err != nil {
			return nil, err
		}

		ref, err := webref.Store(ctx, v.getStore(), o2)
		if err != nil {
			return nil, err
		}

		next := models.Volume{
			ObjectRef: ref,
			Options:   cur.Options,
		}
		return &next, nil
	})
}

func (v *Volume) apply(ctx context.Context, f VolumeMutator) error {
	const maxRetries = 10
	success := false
	for i := 0; !success && i < maxRetries; i++ {
		cur, err := v.cell.Get(ctx)
		if err != nil {
			return err
		}
		curV := models.Volume{}
		if len(cur) > 0 {
			if err := webref.Decode(VolumeCodec, cur, &curV); err != nil {
				return err
			}
		}

		nextV, err := f(curV)
		if err != nil {
			return err
		}
		next, err := webref.Encode(VolumeCodec, nextV)
		if err != nil {
			return err
		}
		success, err = v.cell.CAS(ctx, cur, next)
		if err != nil {
			return err
		}
	}

	if !success {
		return errors.New("could not complete CAS")
	}

	return nil
}

func (v *Volume) Path() Path {
	if v.parent == nil {
		return Path{}
	}
	return v.parent.Path()
}

func (v *Volume) RefIter(ctx context.Context, f func(Ref) bool) (bool, error) {
	cc, err := v.Get(ctx)
	if err != nil {
		return false, err
	}
	cont := f(*cc.ObjectRef)
	return cont, nil
}

func (v *Volume) String() string {
	return "Volume::" + v.cell.ID()
}

func (v *Volume) Size() uint64 {
	return 0
}

func (v *Volume) getStore() ReadWriteOnce {
	fs := v.getFS()
	opts := v.getOptions()
	return &Store{opts: *opts.DataOpts, ms: fs.store}
}

func (v *Volume) getOptions() *Options {
	if v.opts != nil {
		return v.opts
	}
	return DefaultOptions()
}
