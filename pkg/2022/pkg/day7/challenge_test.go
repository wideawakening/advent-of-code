package day7

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge - utils", func() {
	When("given a command", func() {
		It("it determines its a command", func() {

			Expect(IsCommand("$ cd /")).To(Equal(true))
			Expect(IsCommand("dir a")).To(Equal(false))
		})

		It("extracts command and argument if exist", func() {
			command, argument := ExtractCommand("$ cd /")
			Expect(command).To(Equal("cd"))
			Expect(argument).To(Equal("/"))
		})

		It("extract only the command if no arguments", func() {
			command, argument := ExtractCommand("$ ls")
			Expect(command).To(Equal("ls"))
			Expect(argument).To(Equal(""))
		})
	})

	When("cd", func() {
		It("enters a folder when $ cd folder", func() {
			Expect(ProcessCwd("", "/")).To(Equal("/"))
			Expect(ProcessCwd("/", "X")).To(Equal("/X"))
			Expect(ProcessCwd("/X", "Y")).To(Equal("/X/Y"))
		})

		It("leaves the folder when $ cd ..", func() {
			Expect(ProcessCwd("/X/Y", "..")).To(Equal("/X"))
			Expect(ProcessCwd("/", "..")).To(Equal("/"))
		})
	})

	When("ls", func() {
		It("returns expected size", func() {
			Expect(ProcessFileSize("4060174 j")).To(Equal(4060174))
		})

	})

	When("getting subfolders from a path", func() {
		It("returns all paths", func() {
			Expect(GetSubFolders("a/")).To(ContainElements())
			Expect(GetSubFolders("/")).To(ContainElements("/"))
			Expect(GetSubFolders("/a/")).To(ContainElements("/", "/a"))
			Expect(GetSubFolders("/a/b/c/d")).To(ContainElements("/", "/a", "/a/b", "/a/b/c", "/a/b/c/d"))
		})
	})
})

var _ = Describe("challenge - star1", func() {
	When("processed input file", func() {
		It("returns lower than 10K of size directories total size", func() {
			Expect(Star1("input.txt")).To(Equal(1490523))
		})

		It("returns the total size of the directory to delete to be able to install the update", func() {
			Expect(Star2("input.txt")).To(Equal(12390492))
		})
	})
})
