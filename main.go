package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	childPath                 = "skill"
	mainPath                  = "/"
	childPathCore             = "coreComp"
	childPathCourses          = "courses"
	childPathResponsibilities = "responsibilities"
)

// desired json structure in firebase:
// skills/[{recordId1: "skill1", recordId2: "skill2}]

type Skill struct {
	Skill string `json:"skill,omitempty"`
}

type CoreComp struct {
	CoreComp string `json:"coreComp,omitempty"`
}

type Courses struct {
	CourseTitle    string `json:"courseTitle,omitempty"`
	CourseCategory string `json:"courseCategory,omitempty"`
	Completed      string `json:"completed,omitempty"`
	DatePeriod     string `json:"datePeriod,omitempty"`
	Institution    string `json:"institution,omitempty"`
	Description    string `json:"description,omitempty"`
}

type Responsibilities struct {
	Responsibility string `json:"responsibility,omitempty"`
}

var (
	mySkills = []Skill{
		{"Data compilation and manipulation for management reports"},
		{"Detailed project cost estimation"},
		{"Budget compilation"},
		{"Management and cleaning of very large"},
		{"record sets"},
		{"Contractor management"},
		{"Function out sourcing"},
		{"Wireless Local Loop planning"},
		{"Copper line planning"},
		{"Fiber planning"},
		{"Creative GIS system use"},
		{"Right-of-way managemen"},
		{"Site management"},
		{"Safety management"},
	}

	myCoreComp = []CoreComp{
		{"Management of technology"},
		{"GIS project management"},
		{"Project coordination of large projects"},
		{"Business and people leadership"},
		{"Staff management"},
		{"Financial management"},
		{"Business continuity management"},
		{"Health and safety management"},
	}

	myCourses = []Courses{
		{
			CourseTitle:    "Technology Leadership Program",
			CourseCategory: "ManagementLeadership",
			Completed:      "yes",
			DatePeriod:     "1994",
			Institution:    "Stratec",
			Description:    "A yearlong part time program on management of engineering skills throughout the engineering spectrum of South Africa. Thesis subject was - the Internet. ",
		},
		{
			CourseTitle:    "MS Power Point",
			CourseCategory: "Computer literacy",
			Completed:      "yes",
			DatePeriod:     "1996",
			Institution:    "Deloitte & Touche",
			Description:    "Advanced presentation course.",
		},
		{
			CourseTitle:    "Basic and End User - Groupwise: Basic",
			CourseCategory: "Computer literacy",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Emailing system user training.",
		},
		{
			CourseTitle:    "BPM - Introduction to the System, Intro to Simplified Processes",
			CourseCategory: "Project Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house project management system training.",
		},
		{
			CourseTitle:    "Project Scheduling Tool (PST) - PNR and Estimator",
			CourseCategory: "Project Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house project management system training.",
		},
		{
			CourseTitle:    "Effective handling of Discipline",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Labour Relations Training.",
		},
		{
			CourseTitle:    "Labor Relations Kit",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Labour Relations Training.",
		},
		{
			CourseTitle:    "Employee Relations for Managers",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Labour Relations Training.",
		},
		{
			CourseTitle:    "Effective handling of discipline - Engaging the Manager",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Labour Relations Training.",
		},
		{
			CourseTitle:    "Total Quality Management",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "ISO 9001 based qulity training.",
		},
		{
			CourseTitle:    "Management & Administrative Procedures",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "Performance evaluation: Technical - Coaching",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "Full Range Leadership (The Leading Manager - The Coaching Manager - The Empowering Manager - The Team Enabling Manager)",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "What Every Employee Should Know About Human Resourses",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "Customer Relationship: Ambassador Program",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "Ethics",
			CourseCategory: "Management and people skills",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "General People Managent Training.",
		},
		{
			CourseTitle:    "Capital Investment Decision Making Tool",
			CourseCategory: "Finacial Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house Financial and Investment Decision Making Training.",
		},
		{
			CourseTitle:    "Introduction to BCM for Management Level",
			CourseCategory: "Finacial Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house Financial and Investment Decision Making Training.",
		},
		{
			CourseTitle:    "Management PCA evaluation",
			CourseCategory: "Finacial Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house Financial and Investment Decision Making Training.",
		},
		{
			CourseTitle:    "Financial Management for Line Managers",
			CourseCategory: "Finacial Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "In house Financial and Investment Decision Making Training.",
		},
		{
			CourseTitle:    "H&S Seminar (HASSEM) - Management of Occupational Safety: MOOS - SHE for Managers and Supervisors  ",
			CourseCategory: "Saftey Management",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Occupational Saftey Training.",
		},
		{
			CourseTitle:    "E10 introduction",
			CourseCategory: "Technical Training",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Digital Exchange Switching Systems training.",
		},
		{
			CourseTitle:    "EWSD Introduction & Systems Course",
			CourseCategory: "Technical Training",
			Completed:      "yes",
			DatePeriod:     "1996 – 2014",
			Institution:    "Telkom SA",
			Description:    "Digital Exchange Switching Systems training.",
		},
		{
			CourseTitle:    "UNIGIS",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "1999",
			Institution:    "PE Technicon,",
			Description:    "UNIGIS Post Graduate Diploma Technical Workshop Module.",
		},
		{
			CourseTitle:    "Unix for Users",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "2000",
			Institution:    "GIMS",
			Description:    "UNIX foudational course for GIS sytems.",
		},
		{
			CourseTitle:    "Administration of GEO Database",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "2000",
			Institution:    "GIMS",
			Description:    "GIS Adminsiration of Spacial data.",
		},
		{
			CourseTitle:    "ARCINFO 8",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "2000",
			Institution:    "GIMS",
			Description:    "ESRI ARCINFO Systems course.",
		},
		{
			CourseTitle:    "ARCGIS related courses",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "2006",
			Institution:    "GIMS",
			Description:    "ESRI ARCGIS Systems course bridge.",
		},
		{
			CourseTitle:    "GE Small World Systems Course",
			CourseCategory: "GIS ",
			Completed:      "yes",
			DatePeriod:     "2001",
			Institution:    "IST",
			Description:    "IST Small World training.",
		},
		{
			CourseTitle:    "Intro to Java2 programming",
			CourseCategory: "ProgrammingDeveloper",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "INCUSDATA",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Go (Golang) The complete bootcamp.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Learn how to code Google’s Go(Golang) Programming Language.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Web Development with Google’s Go(Golang) Programming Language.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Master Go (Golang) Programming: The Complete Go Bootcamp 2020.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Web Authentication, Encryption, JWT, HMAC, & OAuth With Go.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Mastering Go Modules, gRPC, Crawling, and Collaboration with Git",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Complete Guide to Protocol Buffers 3 [Java, Golang, Python].",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "gRPC [Golang] Master Class: Build Modern API & Microservices.",
			CourseCategory: "GolangDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "The Complete Python bootcamp.",
			CourseCategory: "PythonDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "100 Python exercises: Evaluate and improve your skills.",
			CourseCategory: "PythonDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "NLP – Natural Language Processing with Python.",
			CourseCategory: "PythonDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "The Ultimate MYSQL bootcamp.",
			CourseCategory: "SQLDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Elasticsearch 7 and the Elastic Stack - In Depth & Hands On.",
			CourseCategory: "ElasticSearchDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Ethical Hacking and Network Security from Scratch 2021.",
			CourseCategory: "SecurityDevelopment",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Linux Administration: The Complete Linux Bootcamp 2021.",
			CourseCategory: "LinuxAdministrator",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Agile Crash Course: Agile Project Management; Agile Delivery",
			CourseCategory: "ProjectManagement",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Vue - The Complete Guide (incl. Router & Composition API) ",
			CourseCategory: "FrontEndFramework",
			Completed:      "yes",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "The Web developer bootcamp.",
			CourseCategory: "WebDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Learning Python for Data analysis and visualization.",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "The Python Mega course: Build 10 real world applications.",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Machine Learning A-Z: Hands-on Python and R in Data Science.",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Python for Computer Vision with Open CV and Deep Learning.",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Python & Machine Learning for Financial Analysis. ",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Automate Everything with Python",
			CourseCategory: "PythonDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Modern Artificial Intelligence Masterclass: Build 6 Projects. ",
			CourseCategory: "AIDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Machine Learning & Data Science Foundations Masterclass.   ",
			CourseCategory: "AIDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Learn SQL for Beginners: The Comprehensive Hands-on Bootcamp.",
			CourseCategory: "SQLDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
		{
			CourseTitle:    "Learn SQL for Beginners - How To Learn SQL The Easy Way",
			CourseCategory: "SQLDevelopment",
			Completed:      "in progress",
			DatePeriod:     "2018-currently",
			Institution:    "UDEMY Online",
			Description:    "Programming course.",
		},
	}

	myResponsililities = []Responsibilities{
		{"Planning of technology"},
		{"Team leadership"},
		{"Staff management"},
		{"Financial management"},
		{"GIS management"},
		{"Right-of-way management"},
		{"Record keeping system management"},
		{"Telecommunication site management"},
		{"Business continuity management"},
		{"Health and safety management"},
		{"Contract management"},
		{"Program management"},
	}
)

func main() {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://my-cv-7674b-default-rtdb.firebaseio.com",
		ProjectID:   "my-cv",
	}
	// Initialize the Firebase SDK

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("/Users/hendrikpienaar/Documents/keys/Private/Firebase/my-cv-7674b-firebase-adminsdk-h1aoi-2426d03f74.json")

	// Initialize the app with a service account, granting admin privileges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app:", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	// create the ref
	ref := client.NewRef(mainPath)
	// create the child node
	childRef := ref.Child(childPathResponsibilities)
	// childRef, err := newRef.Push(ctx, nil)
	// if err != nil {
	// 	log.Println("Error pushing child node:", err)
	// }

	// Write the skills to firebase
	// for i := 0; i < len(mySkills); i++ {
	// 	data := mySkills[i]
	// 	err := writeNode(ctx, childRef, data)
	// 	if err != nil {
	// 		log.Printf("Error at writing to db at: %d\n", i)
	// 		break
	// 	}
	// }
	// Write the competencies to firebase
	// for i := 0; i < len(myCoreComp); i++ {
	// 	data := myCoreComp[i]
	// 	err := writeNode(ctx, childRef, data)
	// 	if err != nil {
	// 		log.Printf("Error at writing to db at: %d\n", i)
	// 		break
	// 	}
	// }
	// Write the coourses to firebase
	// for i := 0; i < len(myCourses); i++ {
	// 	data := myCourses[i]
	// 	err := writeNode(ctx, childRef, data)
	// 	if err != nil {
	// 		log.Printf("Error at writing to db at: %d\n", i)
	// 		break
	// 	}
	// }
	// Write the responsibilities to firebase
	for i := 0; i < len(myResponsililities); i++ {
		data := myResponsililities[i]
		err := writeNode(ctx, childRef, data)
		if err != nil {
			log.Printf("Error at writing to db at: %d\n", i)
			break
		}
	}

	// As an admin, the app has access to read and write all data, regardless of Security Rules
	// Read all the data at ref

	// var data map[string]interface{}

	// if err := ref.Get(ctx, &data); err != nil {
	// 	log.Fatalln("Error reading from database:", err)
	// }
	// fmt.Println("Data from database:", data)

}

// Write the skills to the skills node
func writeNode(ctx context.Context, ref *db.Ref, data interface{}) error {
	if _, err := ref.Push(ctx, &data); err != nil {
		log.Printf("Error pushing %v: %v\n", data, err)
		return err
	}
	return nil
}
