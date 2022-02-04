package main

import (
	"log"
	"os"
	"path/filepath"

	"renamer/analyze"
	"renamer/data"
	"renamer/edit"
	"renamer/output"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func main() {
	app := &cli.App{
		Name:            "Renamer",
		Usage:           Version,
		Description:     "音声合成ソフトで生成されたテキストファイルからファイル名をリネームします\nついでにテキスト内容もいい感じに書き換えます",
		Version:         Version,
		HideHelpCommand: true,
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:     "text",
				Aliases:  []string{"t"},
				Required: true,
				Usage:    "リネーム元のテキストファイルのパス",
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	textpath := c.String("text")
	exep, err := os.Executable()
	if err != nil {
		return err
	}

	stg, err := ReadSettings(filepath.Join(filepath.Dir(exep), "setting.yml"))
	if err != nil {
		return err
	}

	//テキストファイル名とテキストファイルのペア作成
	tgts, err := analyze.GetPtnTgts(textpath)
	if err != nil {
		return err
	}

	//パターンからname,body獲得
	tinfo, err := analyze.GetTalkInfo(tgts, stg.Ptns)
	if err != nil {
		return err
	}

	editor := edit.Editor{
		Settings: *stg,
	}

	//置換
	tinfo, err = editor.ReplaceName(*tinfo)
	if err != nil {
		return err
	}

	//ファイル内容修正
	err = output.FixTextFile(textpath, editor.FixBody(tinfo.Body), stg.Rslt.Enc)
	if err != nil {
		return err
	}

	//ファイルリネーム
	newname, err := editor.ReplaceResult(tinfo)
	if err != nil {
		return err
	}
	newname = editor.GetLimitedFileName(newname)

	rnpaths := append(c.Args().Slice(), textpath)
	err = output.Renames(newname, rnpaths)

	if err != nil {
		return err
	}

	return nil
}

func ReadSettings(path string) (*data.Settings, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	r := data.Settings{}
	err = yaml.Unmarshal(b, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
