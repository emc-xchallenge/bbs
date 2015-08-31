package consul_test

import (
	"time"

	"github.com/pivotal-golang/lager/lagertest"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/ginkgomon"

	"github.com/cloudfoundry-incubator/bbs/db"
	"github.com/cloudfoundry-incubator/bbs/models"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/services_bbs"
	oldmodels "github.com/cloudfoundry-incubator/runtime-schema/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/clock/fakeclock"
)

var _ = Describe("CellsLoader", func() {
	Describe("Cells", func() {

		const ttl = 10 * time.Second
		const retryInterval = time.Second
		var (
			clock *fakeclock.FakeClock

			bbs                *services_bbs.ServicesBBS
			presence1          ifrit.Process
			presence2          ifrit.Process
			firstCellPresence  oldmodels.CellPresence
			secondCellPresence oldmodels.CellPresence
			logger             *lagertest.TestLogger
		)

		BeforeEach(func() {
			logger = lagertest.NewTestLogger("test")
			clock = fakeclock.NewFakeClock(time.Now())
			bbs = services_bbs.New(consulSession, clock, logger)

			firstCellPresence = oldmodels.NewCellPresence("first-rep", "1.2.3.4", "the-zone", oldmodels.NewCellCapacity(128, 1024, 3), []string{}, []string{})
			secondCellPresence = oldmodels.NewCellPresence("second-rep", "4.5.6.7", "the-zone", oldmodels.NewCellCapacity(128, 1024, 3), []string{}, []string{})

			presence1 = nil
			presence2 = nil

		})

		AfterEach(func() {
			ginkgomon.Interrupt(presence1)
			ginkgomon.Interrupt(presence2)
		})

		Context("when there is a single cell", func() {
			var cellsLoader db.CellsLoader
			var cells models.CellSet
			var err error

			BeforeEach(func() {
				cellsLoader = consulDB.NewCellsLoader(logger)
				presence1 = ifrit.Invoke(bbs.NewCellPresence(firstCellPresence, retryInterval))

				Eventually(func() ([]oldmodels.CellPresence, error) {
					return bbs.Cells()
				}).Should(HaveLen(1))

				cells, err = cellsLoader.Cells()
			})

			It("returns only one cell", func() {
				Expect(err).NotTo(HaveOccurred())
				Expect(cells).To(HaveLen(1))
				Expect(cells).To(HaveKey("first-rep"))
			})

			Context("when one more cell is added", func() {
				BeforeEach(func() {
					presence2 = ifrit.Invoke(bbs.NewCellPresence(secondCellPresence, retryInterval))

					Eventually(func() ([]oldmodels.CellPresence, error) {
						return bbs.Cells()
					}).Should(HaveLen(2))
				})

				It("returns only one cell", func() {
					cells, err := cellsLoader.Cells()
					Expect(err).NotTo(HaveOccurred())
					Expect(cells).To(HaveLen(1))
				})

				Context("when a new loader is created", func() {
					It("returns two cells", func() {
						newCellsLoader := consulDB.NewCellsLoader(logger)
						cells, err := newCellsLoader.Cells()
						Expect(err).NotTo(HaveOccurred())
						Expect(cells).To(HaveLen(2))
						Expect(cells).To(HaveKey("first-rep"))
						Expect(cells).To(HaveKey("second-rep"))
					})
				})
			})
		})
	})
})