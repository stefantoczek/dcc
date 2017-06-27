package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("v1tov32", func() {
	Describe("parseRestartPolicy", func() {

		tests := map[string]string{
			"always":         "",
			"unless-stopped": "",
			"no":             "none",
			"other":          "other",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(parseRestartPolicy(t)).To(Equal(e))
			}
		})
	})

	Describe("getProto", func() {
		tests := map[string]string{
			"":    "tcp",
			"udp": "udp",
		}

		It("returns proper values", func() {
			for t, e := range tests {
				Expect(getProto(t)).To(Equal(e))
			}
		})
	})

	Describe("PortRange validate", func() {
		Context("invalid port range", func() {
			pr := PortRange{
				Start: 100,
				End:   50,
			}

			f := func() {
				pr.validate("test")
			}

			It("panics", func() {
				Expect(f).To(Panic())
			})
		})

		Context("valid port range", func() {
			pr := PortRange{
				Start: 100,
				End:   150,
			}

			f := func() { pr.validate("test") }

			It("does not panic", func() {
				Expect(f).NotTo(Panic())
			})
		})
	})

	Describe("extractRegexpVars", func() {
		re := createRegexp(`^(?P<volume_name>[a-z0-9-_]+):(?P<container_path>.+)$`)

		str := "test:/var/www"

		It("expects to match the test string", func() {
			Expect(re.MatchString(str)).To(BeTrue())
		})

		exp := map[string]string{
			"volume_name":    "test",
			"container_path": "/var/www",
		}

		It("expects to properly extract the vars", func() {
			Expect(extractRegexpVars(re, &str)).To(Equal(exp))
		})
	})
})
