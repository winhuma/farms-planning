package myconnect

// import (
// 	"context"
// 	"farms-planning/pkg/myfunc"
// 	"fmt"
// 	"os"

// 	"cloud.google.com/go/firestore"
// 	"cloud.google.com/go/storage"
// 	firebase "firebase.google.com/go"
// 	"github.com/rs/zerolog/log"
// 	"google.golang.org/api/option"
// )

// var FIREBASE *firebase.App

// func FirebaseConnect() *firebase.App {
// 	opt := option.WithCredentialsJSON([]byte(os.Getenv("FIREBASE_SDK")))
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Fatal().Msg(myfunc.MyErrFormat(err).Error())
// 	}
// 	FIREBASE = app
// 	return app
// }

// func FirebaseGetStorage() *storage.BucketHandle {
// 	storage, err := FIREBASE.Storage(context.TODO())
// 	if err != nil {
// 		log.Fatal().Msg(myfunc.MyErrFormat(err).Error())
// 	}
// 	bucket, err := storage.Bucket(os.Getenv("FIREBASE_BUCKET"))
// 	if err != nil {
// 		log.Fatal().Msg(myfunc.MyErrFormat(err).Error())
// 	}
// 	return bucket
// }

// func FirebaseGetDB() *firestore.Client {
// 	myfirestore, err := FIREBASE.Firestore(context.TODO())
// 	if err != nil {
// 		log.Fatal().Msg(myfunc.MyErrFormat(err).Error())
// 	}
// 	return myfirestore
// }

// func GetUrlFilePath(filename string) string {
// 	return fmt.Sprintf(
// 		"https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media",
// 		os.Getenv("FIREBASE_BUCKET"),
// 		filename,
// 	)
// }

// import (
// 	"context"
// 	"plan-farm/myconfig/myvar"
// 	"plan-farm/pkg/myconnect"
// 	"plan-farm/pkg/myfunc"

// 	"google.golang.org/api/iterator"
// )

// func FirebaseGetByCollection(CollectionName string) ([]map[string]interface{}, error) {
// 	var mydata = []map[string]interface{}{}
// 	fs := myconnect.FirebaseGetDB()
// 	iter := fs.Collection(CollectionName).Documents(context.TODO())
// 	defer iter.Stop()
// 	for {
// 		doc, err := iter.Next()
// 		if err == iterator.Done {
// 			break
// 		}
// 		if err != nil {
// 			return nil, myfunc.MyErrFormat(err)
// 		}
// 		mydata = append(mydata, doc.Data())
// 	}
// 	return mydata, nil
// }

// func FBGetWaterRequireDay() ([]map[string]interface{}, error) {
// 	areaData, err := FirebaseGetByCollection(myvar.CollectionWaterArea)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return areaData, nil
// }

// func FBGetWaterRequireIndustry() ([]map[string]interface{}, error) {
// 	industryData, err := FirebaseGetByCollection(myvar.CollectionWaterIndustry)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return industryData, nil
// }
