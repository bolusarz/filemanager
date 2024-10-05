package cmd

import (
	"FileOrganizer/models"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var store = models.NewFileDataStore("")

var verbose bool
var extensions []string
var folderName string

var rootCmd = &cobra.Command{
	Use:     "org",
	Short:   "Org is a file organizer",
	Long:    `Org is a file organizer that categorizes and moves files into folders based on their extensions.`,
	Version: "0.1",
}

var categoryListCmd = &cobra.Command{
	Use:   "category",
	Short: "Lists all file categories",
	Long:  "Displays a list of all categories used for organizing files, optionally with extensions.",
	Example: `org category 
org category --verbose`,
	Run: func(cmd *cobra.Command, args []string) {
		DisplayCategories(verbose)
	},
}

var addCategoryCmd = &cobra.Command{
	Use:   "add [category]",
	Short: "Adds a new file category",
	Long: `Creates a new file category with a folder name and a list of associated file extensions.
				You can also remove pre-existing extensions by prepending a dash. The flags are also optional`,
	Example: ` 
		org add documents --ext ".pdf,.docx" --folder "Documents" 
		To Remove Extensions: org add documents --ext "-.pdf"
			`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a category")
		}
		if strings.TrimSpace(args[0]) == "" {
			return errors.New("you must provide a category")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		category := strings.TrimSpace(args[0])
		var validExtensions []string

		for _, extension := range extensions {
			validExtensions = append(
				validExtensions,
				strings.Split(strings.ReplaceAll(extension, " ", ""), ",")...,
			)
		}
		addCategory(category, folderName, validExtensions)
	},
}

var removeCategoryCmd = &cobra.Command{
	Use:     "remove [category]",
	Short:   "Removes a file category",
	Long:    `Removes an existing file category from the list of available categories.`,
	Example: `org remove documents`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a category")
		}
		if strings.TrimSpace(args[0]) == "" {
			return errors.New("you must provide a category")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		category := strings.TrimSpace(args[0])
		removeCategory(category)
	},
}

var organizeCmd = &cobra.Command{
	Use:     "organize [directory]",
	Short:   "Organizes files in a directory",
	Long:    `Organizes files in the specified directory by moving them into folders based on their categories and file extensions.`,
	Example: "org organize /path/to/directory",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("you must provide a directory to organize")
		}
		if strings.TrimSpace(args[0]) == "" {
			return errors.New("you must provide a directory to organize")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		organizeFiles(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	categoryListCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Displays the extensions associated with each category")
	addCategoryCmd.Flags().StringVarP(&folderName, "folder", "f", "", "Folder name for new category")
	addCategoryCmd.Flags().StringArrayVarP(&extensions, "ext", "e", []string{}, "List of extensions for new category")

	rootCmd.AddCommand(categoryListCmd, addCategoryCmd, removeCategoryCmd, organizeCmd)
}
