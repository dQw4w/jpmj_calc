package calc

import (
	"errors"
	"net/http"

	service "github.com/dQw4w/jpmj_calc/services/calc"
	"github.com/dQw4w/jpmj_calc/services/calc/combination"
	"github.com/dQw4w/jpmj_calc/services/calc/win"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type calcRequest struct {
	Hand     string   `form:"hand" binding:"required"`
	Akarehai string   `form:"akarehai" binding:"required"`
	Furolist []string `form:"furo"`

	SelfWind  win.Wind `form:"self_wind" binding:"required"`
	FieldWind win.Wind `form:"field_wind" binding:"required"`

	Tsumo   bool `form:"tsumo" `
	Menchin bool `form:"menchin" `

	Reach       bool `form:"reach"`
	DoubleReach bool `form:"double_reach"`
	ChanKan     bool `form:"chankan"`
	RinShan     bool `form:"rinshan"`
	HaiTei      bool `form:"haitei"`
	HoTei       bool `form:"hotei"`
	Ippatsu     bool `form:"ippatsu"`

	TenHo bool `form:"tenho"`
	JiHo  bool `form:"jiho"`

	MotedoraSuit []string `form:"motedora_suit"`
	MotedoraRank []uint8  `form:"motedora_rank"`

	UradoraSuit []string `form:"uradora_suit"`
	UradoraRank []uint8  `form:"uradora_rank"`
}

// omiempty

func (req calcRequest) ToCommonWin() (win.Common_Win, error) {
	moteSuit, err := StringArrToByteArr(req.MotedoraSuit)
	if err != nil {
		log.Error(err)
		return win.Common_Win{}, err
	}
	uraSuit, err := StringArrToByteArr(req.UradoraSuit)
	if err != nil {
		log.Error(err)
		return win.Common_Win{}, err
	}
	commonWin := win.Common_Win{
		MenziList:    [4]combination.Menzi{},
		Eye:          combination.Pair{},
		Tsumo:        req.Tsumo,
		Menchin:      req.Menchin,
		SelfWind:     uint8(req.SelfWind), //TODO:
		FieldWind:    uint8(req.FieldWind),
		Reach:        req.Reach,
		DoubleReach:  req.DoubleReach,
		ChanKan:      req.ChanKan,
		RinShan:      req.RinShan,
		HaiTei:       req.HaiTei,
		HoTei:        req.HoTei,
		Ippatsu:      req.Ippatsu,
		TenHo:        req.TenHo,
		JiHo:         req.JiHo,
		MotedoraSuit: moteSuit,
		MotedoraRank: req.MotedoraRank,
		UradoraSuit:  uraSuit,
		UradoraRank:  req.UradoraRank,
	}

	return commonWin, nil
}

func (ctl *CalculateController) CalculateResults(c *gin.Context) {

	var request calcRequest
	err := c.ShouldBind(&request)

	// c.JSON(http.StatusAccepted, request)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	var otherArgs win.Common_Win
	otherArgs, err = request.ToCommonWin()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	// fmt.Print("s")
	log.Print(otherArgs)
	s := service.NewCalculateService()
	c.String(http.StatusOK, s.Calculate(request.Hand, request.Akarehai, request.Furolist, otherArgs))
}

func StringArrToByteArr(strs []string) ([]byte, error) {

	var output []byte
	log.Warn("output:\n", output)
	for _, str := range strs {
		if len(str) > 1 {
			return nil, errors.New("Invalid")
		}
		output = append(output, byte(str[0]))
	}
	return output, nil
}
