package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

const (
	childPath                  = "skill"
	mainPath                   = "/"
	childPathCore              = "coreComp"
	childPathCourses           = "courses"
	childPathResponsibilities  = "responsibilities"
	childPathEmploymentHistory = "employmentsHistory"
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

type EmploymentHistory struct {
	JobTitle             string   `json:"jobTitle,omitempty"`
	Place                string   `json:"place,omitempty"`
	DurationPeriod       string   `json:"durationPeriod,omitempty"`
	JobDescription1      string   `json:"jobDescription1,omitempty"`
	JobDescription2      string   `json:"jobDescription2,omitempty"`
	JobDescription3      string   `json:"jobDescription3,omitempty"`
	JobResponsibilities1 []string `json:"jobResponsibilities1,omitempty"`
	JobResponsibilities2 []string `json:"jobResponsibilities2,omitempty"`
	JobResponsibilities3 []string `json:"jobResponsibilities3,omitempty"`
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

	newCourse = Courses{
		CourseTitle:    "Mastering Go Programming",
		CourseCategory: "Golang/Development",
		Completed:      "yes",
		DatePeriod:     "2023-currently",
		Institution:    "UDEMY Online",
		Description:    "Programming course",
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

	myEmploymentHistory = []EmploymentHistory{
		{
			JobTitle:       `Manager: Access Engineering, Fiber allocation and Engineering Operations, Telkom SA`,
			Place:          "Port Elizabeth, South Africa",
			DurationPeriod: "August 2012 to September 2014",
			JobDescription1: `Managed 75 staff members reporting to eight supervisors, spread over the entire region in Port Elizabeth, George and East London. This included seven diverse work disciplines (access lines planning, regional operations and maintenance co-ordination of cable pair replacement recording, capturing of infrastructure records, GIS regional functions, regional right-of-way functions, radio site management and fiber pair allocations.)
			My 2014 Operating budget was mR7.9 and the Capital budget (mR49 CAPEX, mR10 OPEX) to facilitate the financial year’s major electronic exchange conversion project to street cabinet exchange units (Fiber to the Cabinet - FTTC) (914 individual projects).
			Special tasks: Development and maintenance of business continuity plans for the Telkom Southern Region (Southern Cape, Eastern Cape and Border areas).
			`,
		},
		{
			JobTitle:        "Manager: Access Engineering and Engineering Operations, Telkom SA",
			Place:           "Port Elizabeth, South Africa",
			DurationPeriod:  "September 2011 to July 2012",
			JobDescription1: "Management of 40 staff members, sections included",
			JobResponsibilities1: []string{
				"Building cabling and customer premises facility management.",
				"Regional operations and maintenance coordination of cable pair replacement recording.",
				"Capturing of infrastructure records.",
				"GIS regional functions.",
				"Regional right-of-way functions.",
				"Conformance testing.",
				"Material management.",
			},
		},
		{
			JobTitle:       "Manager: Regional Engineering Operations, Telkom SA",
			Place:          "Port Elizabeth, South Africa",
			DurationPeriod: "July 2005 to August 2011",
			JobDescription1: `Management of the engineering operation sections consisted of engineering support sections such as: Right-of-way, GIS and the engineering location records sections, building cabling and customer premises facility management for the Southern Region (Eastern Cape, Southern Cape and Border).  
			As part of the national team responsible for the NETPLAN (GE Smallworld GIS) implementation in Telkom, roles and responsibilities included`,
			JobResponsibilities1: []string{
				"Revision of planning principles and procedures and spatial information requirements.",
				"Guidelines of data capturing and management of regional contractor manned capturing teams and targets.",
			},
			JobDescription2: `Management of Southern region’s portion of NETDATA, a multi-year project, capturing network infrastructure of both physical, logical and equipment connections. 
			Responsibilities included`,
			JobResponsibilities2: []string{
				"Capturing program including staffing models and budgets.",
				"Data capturing targets management.",
				"Quality management of captured data.",
				"Control and management of the sourcing and verification of the record infrastructure data, undertaken by in house and contractor-based teams.",
			},
			JobDescription3: "The Operations section consisted of 38 staff members. Sections managed",
			JobResponsibilities3: []string{
				"Regional operations and maintenance coordination of cable pair replacement recording.",
				"Capturing of infrastructure records sections.",
				"GIS regional functions.",
				"Regional right-of-way functions.",
				"Material management (checking of estimated material with physical material available).",
			},
		},
		{
			JobTitle:             "Manager: Regional Network Engineering, Telkom SA",
			Place:                "Port Elizabeth, South Africa",
			DurationPeriod:       "July 2000 to June 2005",
			JobDescription1:      "Name change, due to restructuring at Telkom with added responsibility for Southern Cape area.",
			JobResponsibilities1: []string{},
		},
		{
			JobTitle:       "Manager: Network Planning, Telkom SA",
			Place:          "Port Elizabeth, South Africa",
			DurationPeriod: "April 1999 to June 2000",
			JobDescription1: `Manager of the Midlands lines planning section responsible for automation of rural farm lines and town planning of fix lines subscribers employing various technologies.
			Responsibilities included`,
			JobResponsibilities1: []string{
				"General management of staff and contractors.",
				"Liaison with ESCOM and local municipalities.",
				"Coordinating and contribution to the compilation of national right-of-way documentation for Telkom",
			},
		},
		{
			JobTitle:       "Manager: Network Planning, Telkom SA",
			Place:          "Port Elizabeth, South Africa",
			DurationPeriod: "August 1992 to March 1999",
			JobDescription1: `Senior Engineer Lines Planning (country and rural areas). Registered as professional engineer in 1993.
			Heading planning teams in rural areas, involving farm line conversions with various technologies, such as, SOR18, RURTEL and point-to-point radio systems. 
			This position evolved to the position of Senior Engineer Wireless Local Loop Planning: As planning team leader and planning project manager, leading teams consisting of in-house planning teams and in sourced planning teams from ESMARTEL, CSIR and ALCATEL, the Wireless conversion and connection of subscribers, spanning the entire Southern Region (greater Eastern Cape, Southern Cape and Ciskei/Transkei area) was planned (108 sites with 20 & 35 meter masts in urban areas 533 and in rural areas).  
			The 1998/99 financial budget to provide 112 000 wireless links was mR350.  The tasks coordinated included: Site identification using GIS tools, WWL planning, detail planning, right-of-way procurement, planning acceptance and quality control, footprint verification.  In the process, special original, innovative tools were developed to enable rural subscribers to apply for service, installation teams to connect, and maintenance teams to provide point to multi point wireless network services.  
			This was recognized by an innovation award from Telkom (see Key Achievements).`,
			JobResponsibilities1: []string{},
		},
		{
			JobTitle:             "Engineer, Telkom SA",
			Place:                "Port Elizabeth, South Africa",
			DurationPeriod:       "",
			JobDescription1:      "",
			JobResponsibilities1: []string{},
		},
		{
			JobTitle:       "Engineer, Telkom SA",
			Place:          "Port Elizabeth, South Africa",
			DurationPeriod: "January 1989 to July 1992",
			JobDescription1: `Engineer EWSD (German Electronic Switches) and later Engineer Electronic Exchanges (EWSD and E10 (French Electronic Switches).  This entailed operations and maintenance of all electronic exchanges in the Eastern Cape region.

			`,
			JobResponsibilities1: []string{},
		},
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
	childRef := ref.Child(childPathCourses)
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
	// for i := 0; i < len(myResponsililities); i++ {
	// 	data := myResponsililities[i]
	// 	err := writeNode(ctx, childRef, data)
	// 	if err != nil {
	// 		log.Printf("Error at writing to db at: %d\n", i)
	// 		break
	// 	}
	// }
	// // Write the responsibilities to firebase
	// for i := 0; i < len(myEmploymentHistory); i++ {
	// 	data := myEmploymentHistory[i]
	// 	err := writeNode(ctx, childRef, data)
	// 	if err != nil {
	// 		log.Printf("Error at writing to db at: %d\n", i)
	// 		break
	// 	}
	// }
	// As an admin, the app has access to read and write all data, regardless of Security Rules
	// Read all the data at ref

	// var data map[string]interface{}

	// if err := ref.Get(ctx, &data); err != nil {
	// 	log.Fatalln("Error reading from database:", err)
	// }
	// fmt.Println("Data from database:", data)

	// Write the new course to firebase
	data := newCourse
	err = writeNode(ctx, childRef, data)
	if err != nil {
		log.Println("Error at writing new course to db")
	}
}

// Write the a course to the course node
func writeNode(ctx context.Context, ref *db.Ref, data interface{}) error {
	if _, err := ref.Push(ctx, &data); err != nil {
		log.Printf("Error pushing %v: %v\n", data, err)
		return err
	}
	return nil
}
