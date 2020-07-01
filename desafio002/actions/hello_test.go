package actions

import "net/http"

func (as *ActionSuite) Test_Hello_Index() {
	res := as.HTML("/hello").Get()

	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Hello Full Cycle")
}
