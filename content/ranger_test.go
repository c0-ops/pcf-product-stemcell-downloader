// Copyright 2017-Present Pivotal Software, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package content_test

import (
	"github.com/c0-ops/pcf-product-stemcell-downloader/content"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildRange", func() {
	Context("when an even content length is provided", func() {
		var cr content.Ranger

		BeforeEach(func() {
			cr = content.NewRanger(10)
		})

		It("returns an set of byte ranges", func() {
			contentLength := int64(100)
			r, err := cr.BuildRange(contentLength)
			Expect(err).NotTo(HaveOccurred())

			Expect(r).To(Equal([]string{
				"0-9",
				"10-19",
				"20-29",
				"30-39",
				"40-49",
				"50-59",
				"60-69",
				"70-79",
				"80-89",
				"90-99",
			}))
		})
	})

	Context("when an odd content length is provided", func() {
		var cr content.Ranger

		BeforeEach(func() {
			cr = content.NewRanger(10)
		})

		It("returns the byte ranges", func() {
			contentLength := int64(101)
			r, err := cr.BuildRange(contentLength)
			Expect(err).NotTo(HaveOccurred())

			Expect(r).To(Equal([]string{
				"0-9",
				"10-19",
				"20-29",
				"30-39",
				"40-49",
				"50-59",
				"60-69",
				"70-79",
				"80-89",
				"90-100",
			}))
		})
	})

	Context("when content length is less than the number of hunks", func() {
		var cr content.Ranger

		BeforeEach(func() {
			cr = content.NewRanger(10)
		})

		It("returns as many byte ranges as possible", func() {
			contentLength := int64(3)
			r, err := cr.BuildRange(contentLength)
			Expect(err).NotTo(HaveOccurred())

			Expect(r).To(Equal([]string{
				"0-2",
			}))
		})

		It("returns as many byte ranges as possible", func() {
			contentLength := int64(9)
			r, err := cr.BuildRange(contentLength)
			Expect(err).NotTo(HaveOccurred())

			Expect(r).To(Equal([]string{
				"0-1",
				"2-3",
				"4-5",
				"6-8",
			}))
		})
	})

	Context("when an error occurs", func() {
		Context("when the content length is zero", func() {
			var cr content.Ranger

			BeforeEach(func() {
				cr = content.NewRanger(10)
			})

			It("returns an error", func() {
				contentLength := int64(0)
				_, err := cr.BuildRange(contentLength)
				Expect(err).To(MatchError("content length cannot be zero"))
			})
		})
	})
})
