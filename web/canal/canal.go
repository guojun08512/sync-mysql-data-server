package canal

import (
	"fmt"
	"time"
	"github.com/CanalClient/canal-go/client"
	protocol "github.com/CanalClient/canal-go/protocol"
	"github.com/golang/protobuf/proto"

	"sync-mysql-data-server/pkg/config"
)
var conf = config.Config

func Init() error {
	connector := client.NewSimpleCanalConnector(conf.GetString("canal.host"), conf.GetInt("canal.port"), "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		return err
	}
	err = connector.Subscribe(".*\\\\..*")
	if err != nil {
		return err
	}

	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			return err
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			ti := conf.GetInt("canal.schedule")
			time.Sleep(time.Duration(ti)* time.Second)
			fmt.Println("===没有数据了===")
			continue
		}

		printEntry(message.Entries)

	}
}

func printEntry(entrys []protocol.Entry) error {
	for _, entry := range entrys {
		if entry.GetEntryType() == protocol.EntryType_TRANSACTIONBEGIN || entry.GetEntryType() == protocol.EntryType_TRANSACTIONEND {
			continue
		}
		rowChange := new(protocol.RowChange)

		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		if err != nil {
			return err
		}
		if rowChange != nil {
			eventType := rowChange.GetEventType()
			header := entry.GetHeader()
			fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

			for _, rowData := range rowChange.GetRowDatas() {
				if eventType == protocol.EventType_DELETE {
					printColumn(rowData.GetBeforeColumns())
				} else if eventType == protocol.EventType_INSERT {
					printColumn(rowData.GetAfterColumns())
				} else {
					fmt.Println("-------> before")
					printColumn(rowData.GetBeforeColumns())
					fmt.Println("-------> after")
					printColumn(rowData.GetAfterColumns())
				}
			}
		}
	}
	return nil
}

func printColumn(columns []*protocol.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}
