package webfscmd

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/brendoncarroll/webfs/pkg/webfsim"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(applyOptsCmd)
	applyOptsCmd.Flags().StringP("file", "f", "", "-f")
	applyOptsCmd.Flags().StringP("vol", "", "", `--vol=ABCDEF_VOL_ID`)

	getOptsCmd.Flags().StringP("vol", "", "", `--vol=ABCDEF_VOL_ID`)
	rootCmd.AddCommand(getOptsCmd)
}

var applyOptsCmd = &cobra.Command{
	Use:   "apply-options",
	Short: "Apply options, as specified in a file, to a volume",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := setupWfs(); err != nil {
			return err
		}
		filePath, err := cmd.Flags().GetString("file")
		if err != nil {
			return err
		}
		data, err := ioutil.ReadFile(filePath)
		if err != nil {
			return err
		}
		opts := &webfsim.Options{}
		if err = decode(data, opts); err != nil {
			return err
		}
		volID, err := cmd.Flags().GetString("vol")
		if err != nil {
			return err
		}
		vol, err := wfs.GetVolume(ctx, volID)
		if err != nil {
			return err
		}

		return vol.SetOptions(ctx, opts)
	},
}

var getOptsCmd = &cobra.Command{
	Use:   "get-options",
	Short: "Write the options for a volume to stdout",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := setupWfs(); err != nil {
			return err
		}
		volID, err := cmd.Flags().GetString("vol")
		if err != nil {
			return err
		}
		vol, err := wfs.GetVolume(ctx, volID)
		if err != nil {
			return err
		}
		c, err := vol.Get(ctx)
		if err != nil {
			return err
		}
		prettyPrint(cmd.OutOrStdout(), c.Options)
		return nil
	},
}

func decode(p []byte, x interface{}) error {
	pm := x.(proto.Message)
	buf := bytes.NewBuffer(p)
	return jsonpb.Unmarshal(buf, pm)
}

func prettyPrint(w io.Writer, x interface{}) {
	pm := x.(proto.Message)
	m := jsonpb.Marshaler{
		Indent: " ",
	}
	if err := m.Marshal(w, pm); err != nil {
		panic(err)
	}
	_, err := w.Write([]byte("\n"))
	if err != nil {
		panic(err)
	}
}
