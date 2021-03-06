package ddevapp_test

import (
	"github.com/drud/ddev/pkg/ddevapp"
	"github.com/drud/ddev/pkg/testcommon"
	asrt "github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

// TestApptypeDetection does a simple test of various filesystem setups to make
// sure the expected apptype is returned.
func TestApptypeDetection(t *testing.T) {
	assert := asrt.New(t)

	fileLocations := map[string]string{
		"drupal7":   "scripts/drupal.sh",
		"drupal8":   "core/scripts/drupal.sh",
		"wordpress": "wp-login.php",
	}

	for expectedType, expectedPath := range fileLocations {
		testDir := testcommon.CreateTmpDir("TestApptype")

		// testcommon.Chdir()() and CleanupDir() checks their own errors (and exit)
		defer testcommon.CleanupDir(testDir)
		defer testcommon.Chdir(testDir)()

		err := os.MkdirAll(filepath.Join(testDir, filepath.Dir(expectedPath)), 0777)
		assert.NoError(err)

		_, err = os.OpenFile(filepath.Join(testDir, expectedPath), os.O_RDONLY|os.O_CREATE, 0666)
		assert.NoError(err)

		app, err := ddevapp.NewApp(testDir, ddevapp.DefaultProviderName)
		assert.NoError(err)

		foundType := app.DetectAppType()
		assert.EqualValues(expectedType, foundType)
	}

}
