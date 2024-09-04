package cmd

import (
	"FileOrganizer/models"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var store = models.New("")

var verbose bool
var extensions []string
var folderName string

var rootCmd = &cobra.Command{
	Use:     "org",
	Short:   "Org is a file organizer",
	Long:    `Org is a file organizer that sorts your files into folders for easy access`,
	Version: "0.1",
	//Run: func(cmd *cobra.Command, args []string) {}
}

var categoryListCmd = &cobra.Command{
	Use: "category",
	Run: func(cmd *cobra.Command, args []string) {
		DisplayCategories(verbose)
	},
}

var addCategoryCmd = &cobra.Command{
	Use:     "add",
	Example: "org add [category]",
	Short:   "Adds a category to the list",
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
		addCategory(category, folderName, extensions)
	},
}

var removeCategoryCmd = &cobra.Command{
	Use:        "remove",
	ArgAliases: []string{"category"},
	Example:    "org remove [category]",
	Short:      "Removes a category",
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

	rootCmd.AddCommand(categoryListCmd, addCategoryCmd, removeCategoryCmd)
}
