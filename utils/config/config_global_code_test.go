package config

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	name        string
	workingDir  string
	want        string
	code        string
	accountID   string
	campaignID  string
	variationID string
	elementID   string
	selector    string
	wantErr     bool
}

var (
	mockAccountID   = "123456"
	mockCampaignID  = "100000"
	mockVariationID = "200000"
	mockElementID   = "300000"
	mockSelector    = "document.querySelector('main')"
)

func TestMain(m *testing.M) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}

	defer os.RemoveAll(currentDir + "/abtasty")

	m.Run()
}

func TestCheckWorkingDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	tests := []TestStruct{
		{
			name:       "ExistingDirectory",
			workingDir: currentDir,
			want:       currentDir,
			wantErr:    false,
		},
		{
			name:       "NonExistingDirectory",
			workingDir: "/path/to/nonexistent/directory",
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckWorkingDirectory(tt.workingDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckWorkingDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckWorkingDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckGlobalCodeDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	tests := []TestStruct{
		{
			name:       "ExistingDirectory",
			workingDir: currentDir,
			want:       currentDir + "/abtasty",
			wantErr:    false,
		},
		{
			name:       "NonExistingDirectory",
			workingDir: "/path/to/nonexistent/directory",
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckGlobalCodeDirectory(tt.workingDir)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckGlobalCodeDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckGlobalCodeDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountGlobalCodeDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []TestStruct{
		{
			name:       "ExistingDirectory",
			workingDir: currentDir,
			code:       "console.log('Hello, World!')", // Content of JavaScript file
			accountID:  mockAccountID,
			want:       currentDir + "/abtasty/" + mockAccountID + "/accountGlobalCode.js",
			wantErr:    false,
		},
		{
			name:       "NonExistingDirectory",
			workingDir: "/path/to/nonexistent/directory",
			code:       "console.log('Hello, World!')", // Content of JavaScript file
			accountID:  mockAccountID,
			want:       "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccountGlobalCodeDirectory(tt.workingDir, tt.accountID, tt.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("AccountGlobalCodeDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("AccountGlobalCodeDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCampaignGlobalCodeDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []TestStruct{
		{
			name:       "ExistingDirectory",
			workingDir: currentDir,
			code:       "console.log('Hello, World!')", // Content of JavaScript file
			accountID:  "123456",
			campaignID: "100000",
			want:       currentDir + "/abtasty/" + mockAccountID + "/" + mockCampaignID + "/campaignGlobalCode.js",
			wantErr:    false,
		},
		{
			name:       "NonExistingDirectory",
			workingDir: "/path/to/nonexistent/directory",
			code:       "console.log('Hello, World!')", // Content of JavaScript file
			accountID:  "123456",
			campaignID: "100000",
			want:       "",
			wantErr:    true,
		},
	}

	for i, tt := range tests {
		if i == 0 {
			t.Run(tt.name, func(t *testing.T) {
				got, err := CampaignGlobalCodeDirectory(tt.workingDir, tt.accountID, tt.campaignID, tt.code)
				if (err != nil) != tt.wantErr {
					t.Errorf("CampaignGlobalCodeDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got != tt.want {
					t.Errorf("CampaignGlobalCodeDirectory() = %v, want %v", got, tt.want)
				}
			})

		}
	}
}

func TestVariationGlobalCodeDirectoryJS(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []TestStruct{
		{
			name:        "ExistingDirectory",
			workingDir:  currentDir,
			code:        "console.log('Hello, World!')", // Content of JavaScript file
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			want:        currentDir + "/abtasty/" + mockAccountID + "/" + mockCampaignID + "/" + mockVariationID + "/variationGlobalCode.js",
			wantErr:     false,
		},
		{
			name:        "NonExistingDirectory",
			workingDir:  "/path/to/nonexistent/directory",
			code:        "console.log('Hello, World!')", // Content of JavaScript file
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			want:        "",
			wantErr:     true,
		},
	}

	for i, tt := range tests {
		if i == 0 {
			t.Run(tt.name, func(t *testing.T) {
				got, err := VariationGlobalCodeDirectoryJS(tt.workingDir, tt.accountID, tt.campaignID, tt.variationID, tt.code)
				if (err != nil) != tt.wantErr {
					t.Errorf("VariationGlobalCodeDirectoryJS() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got != tt.want {
					t.Errorf("VariationGlobalCodeDirectoryJS() = %v, want %v", got, tt.want)
				}
			})

		}
	}
}

func TestVariationGlobalCodeDirectoryCSS(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []TestStruct{
		{
			name:        "ExistingDirectory",
			workingDir:  currentDir,
			code:        ".id{ \"color\" : black}",
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			want:        currentDir + "/abtasty/" + mockAccountID + "/" + mockCampaignID + "/" + mockVariationID + "/variationGlobalCode.css",
			wantErr:     false,
		},
		{
			name:        "NonExistingDirectory",
			workingDir:  "/path/to/nonexistent/directory",
			code:        ".id{ \"color\" : black}",
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			want:        "",
			wantErr:     true,
		},
	}

	for i, tt := range tests {
		if i == 0 {
			t.Run(tt.name, func(t *testing.T) {
				got, err := VariationGlobalCodeDirectoryCSS(tt.workingDir, tt.accountID, tt.campaignID, tt.variationID, tt.code)
				if (err != nil) != tt.wantErr {
					t.Errorf("VariationGlobalCodeDirectoryCSS() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got != tt.want {
					t.Errorf("VariationGlobalCodeDirectoryCSS() = %v, want %v", got, tt.want)
				}
			})

		}
	}
}

func TestElementModificationCodeDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}

	tests := []TestStruct{
		{
			name:        "ExistingDirectory",
			workingDir:  currentDir,
			code:        "console.log('Hello, World!')",
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			elementID:   mockElementID,
			selector:    mockSelector,
			want:        currentDir + "/abtasty/" + mockAccountID + "/" + mockCampaignID + "/" + mockVariationID + "/" + mockElementID + "/element.js",
			wantErr:     false,
		},
		{
			name:        "NonExistingDirectory",
			workingDir:  "/path/to/nonexistent/directory",
			code:        "console.log('Hello, World!')",
			accountID:   mockAccountID,
			campaignID:  mockCampaignID,
			variationID: mockVariationID,
			elementID:   mockElementID,
			selector:    mockSelector,
			want:        "",
			wantErr:     true,
		},
	}

	for i, tt := range tests {
		if i == 0 {
			t.Run(tt.name, func(t *testing.T) {
				got, err := ElementModificationCodeDirectory(tt.workingDir, tt.accountID, tt.campaignID, tt.variationID, tt.elementID, tt.selector, []byte(tt.code))
				if (err != nil) != tt.wantErr {
					t.Errorf("ElementModificationCodeDirectory() error = %v, wantErr %v", err, tt.wantErr)
				}
				if got != tt.want {
					t.Errorf("ElementModificationCodeDirectory() = %v, want %v", got, tt.want)
				}
			})

		}
	}
}

func TestAddHeaderSelectorComment(t *testing.T) {
	fileCode := AddHeaderSelectorComment("example selector", "console.log('Hello World !')")
	fileContent := []byte("/* Selector: example selector */\nconsole.log('Hello World !')")
	assert.Equal(t, fileContent, fileCode)
}
