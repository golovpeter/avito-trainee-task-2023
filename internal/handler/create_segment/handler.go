package create_segment

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golovpeter/avito-trainee-task-2023/internal/common"
	"github.com/golovpeter/avito-trainee-task-2023/internal/service/create_segment"
)

type handler struct {
	log     *logrus.Logger
	service create_segment.CreateSegmentService
}

func NewHandler(
	log *logrus.Logger,
	service create_segment.CreateSegmentService,
) *handler {
	return &handler{
		log:     log,
		service: service,
	}
}

// Create segment godoc
// @Summary      Create new segment
// @Description	 adding a new segment by name
// @Tags         segments
// @Accept       json
// @Produce      json
// @Param request body CreateSegmentIn true "request"
// @Success 200
// @Failure 400 {object} common.ErrorOut
// @Failure 500 {object} common.ErrorOut
// @Router       /segment/create [post]
func (h *handler) CreateSegment(c *gin.Context) {
	var in CreateSegmentIn

	if err := c.BindJSON(&in); err != nil {
		h.log.WithError(err).Error("error binding JSON")
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: "error binding JSON",
		})
		return
	}

	isValid, errMessage, err := validateInParams(in)

	if err != nil {
		h.log.WithError(err).Error(err.Error())
		c.JSON(http.StatusInternalServerError, common.ErrorOut{
			ErrorMessage: "failed parse regexp",
		})
		return
	}

	if !isValid {
		h.log.Warn(errMessage)
		c.JSON(http.StatusBadRequest, common.ErrorOut{
			ErrorMessage: errMessage,
		})
		return
	}

	err = h.service.CreateSegment(&create_segment.CreateSegmentData{
		SegmentSlug:    in.SegmentSlug,
		PercentOfUsers: in.PercentOfUsers,
	})

	if err != nil {
		h.log.WithError(err).Error("error create segments")
		c.JSON(http.StatusInternalServerError, common.ErrorOut{
			ErrorMessage: "error create segments",
		})
		return
	}
}
