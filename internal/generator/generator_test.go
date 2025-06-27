package generator_test

import (
	"go-scaffold/internal/generator"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerate_GenerateFiles(t *testing.T) {
	// 一時ディレクトリを作成
	tmpDir := t.TempDir()

	// テスト用にworkingディレクトリを設定
	oldDir, _ := os.Getwd()
	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)

	// 出力ファイルの確認
	layers := []string{"handler", "usecase", "service", "repository"}
	// テンプレートをテスト用一時ディレクトリにコピー
	rootDir := filepath.Clean(filepath.Join(oldDir, "..", ".."))
	for _, layer := range layers {
		src := filepath.Join(rootDir, "internal", "template", layer+".tmpl")
		dst := filepath.Join("internal", "template", layer+".tmpl")
		copyTemplate(t, src, dst)
	}

	// テンプレートは、internal/templateディレクトリに配置されていると仮定
	err := generator.Generate("user", []string{"handler", "usecase", "service", "repository"})
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	for _, layer := range layers {
		outFile := filepath.Join("internal", layer, "user.go")
		if _, err := os.Stat(outFile); os.IsNotExist(err) {
			t.Errorf("Expected file %s does not exist", outFile)
		}
		// 中身の確認
		content, _ := os.ReadFile(outFile)
		if !strings.Contains(string(content), "type User struct") {
			t.Errorf("Generated content is missing expected struct definition")
		}
	}
}

// テストテンプレートコピー
func copyTemplate(t *testing.T, src, dst string) {
	t.Helper()
	input, err := os.ReadFile(src)
	if err != nil {
		t.Fatalf("failed to read template %s: %v", src, err)
	}
	err = os.MkdirAll(filepath.Dir(dst), 0755)
	if err != nil {
		t.Fatalf("failed to make dir for template: %v", err)
	}
	err = os.WriteFile(dst, input, 0644)
	if err != nil {
		t.Fatalf("failed to write template to temp dir: %v", err)
	}
}
