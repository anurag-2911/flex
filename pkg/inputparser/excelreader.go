package inputparser

import (
	"assetmgmt/pkg/model"
	"github.com/tealeg/xlsx"
)

func ReadExcel(fileName string) ([]model.Asset, error) {
	var assets []model.Asset
	excelFile, err := xlsx.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	for _, sheet := range excelFile.Sheets {
		for _, row := range sheet.Rows {
			var comp model.Asset
			for i, cell := range row.Cells {
				text := cell.String()
				switch i {
				case 0:
					comp.ComputerID = text
				case 1:
					comp.UserID = text
				case 2:
					comp.ApplicationID = text
				case 3:
					comp.ComputerType = text
				}
			}
			assets = append(assets, comp)
		}
	}
	return assets, nil
}
