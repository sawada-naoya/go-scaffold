package generator

import (
	"fmt"
	"go-scaffold/internal/generator/util"
	"os"
	"path/filepath"
	"text/template"
)

// Generate creates code files for the provided entity using templates for each
// layer. The generated files will be placed under `internal/<layer>/<entity>.go`.
// `name` should be the entity name in snake_case or any format; it will be
// converted to PascalCase for struct names and snake_case for filenames.
func Generate(name string, layers []string) error {
	// 構造体名用に PascalCase に変換（例: user → User）
	structName := util.ToPascalCase(name)
	// ファイル名用に snake_case に変換（例: user → user.go）
	fileName := util.ToSnakeCase(name) + ".go"

	// 各レイヤーに対してコードファイルを生成
	for _, layer := range layers {
		err := generateLayerFile(layer, structName, fileName)
		if err != nil {
			// レイヤー単位でエラーをラップして返却
			return fmt.Errorf("failed to generate file for layer %s: %w", layer, err)
		}
	}
	return nil
}

func generateLayerFile(layer, structName, fileName string) error {
	// 出力ディレクトリのパスを構築（例: internal/handler）
	dir := filepath.Join("internal", layer)
	// テンプレートファイルのパスを構築（例: internal/template/handler.tmpl）
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	// テンプレートファイルのパスを構築（例: internal/template/handler.tmpl）
	tmplPath := filepath.Join("internal", "template", layer+".tmpl")
	// テンプレートファイルをパース
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return err
	}
	// 上書きを避けて一意なファイルパスを取得（user.go / user_copy.go 等）
	filePath, err := util.GetUniqueFilePath(dir, fileName)
	if err != nil {
		return err
	}
	// 出力ファイルを作成
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// テンプレートに埋め込むデータを定義（StructName に置換される）
	data := map[string]string{"StructName": structName}
	// テンプレートを実行してファイルへ書き出し
	if err := tmpl.Execute(f, data); err != nil {
		return err
	}
	fmt.Printf("Created: %s\n", filePath)

	return nil
}
