package challenge

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("challenge", func() {
	Context("Star1", func() {

		It("resolves sample line", func() {
			Expect(GetPossibleArranges("???.### 1,1,3")).To(Equal(1))
			Expect(GetPossibleArranges(".??..??...?##. 1,1,3")).To(Equal(4))

			// no dot
			Expect(GetPossibleArranges("?#?#?#?#?#?#?#? 1,3,1,6")).To(Equal(1))
			//Expect(GetPossibleArranges("????.#...#... 4,1,1")).To(Equal(1))
			//Expect(GetPossibleArranges("????.######..#####. 1,6,5")).To(Equal(4))
			//Expect(GetPossibleArranges("?###???????? 3,2,1")).To(Equal(10))
		})
		XWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(0))
			})
		})
		XWhen("given input", func() {
			It("resolves", func() {
				Expect(Star1("star_input.txt")).To(Equal(0))
			})
		})
	})
	Context("Star2", func() {
		XWhen("given sample", func() {
			It("resolves", func() {
				Expect(Star1("star_sample.txt")).To(Equal(0))
			})
		})
		XWhen("given input", func() {
			It("resolves", func() {
				Expect(Star1("star_input.txt")).To(Equal(0))
			})
		})
	})
})
