## Shilling Scheduling API Go SDK 
The Shilling Scheduling API Go SDK is primarily generated using [go-swagger](https://github.com/go-swagger/go-swagger) which implements a more Go idiomatic (and functioning) version of what the swagger-codegen project does for other languages.  While we generate the client, models, and core requests handling we also add some helper utilities to symplify instantiation.  Below are some basic examples to get you up and running in short order.  

### Prerequisites 
You will need to request an environment and access credentials from our site [Shilling](https://www.shilling.io/scheduling).  After which you should be able to fill in the necessary values below in the examples.

### Install
This package uses Go modules so to install you can simply run the following
```
go get github.com/shillingio/scheduling.go
```

### Usage

```go
  import(
    "github.com/go-openapi/swag"

    "github.com/shillingio/scheduling.go"
    "github.com/shillingio/scheduling.go/models"
    "github.com/shillingio/scheduling.go/scheduler"
  )

  api, err := scheduling.NewAPI(scheduling.Config{
		BaseURL:      "<provided base URL>",
		TokenURL:     "<provided token URL>",
		ClientID:     "<provided client id>",
		ClientSecret: "<provided client secret>",
	})
```

##### Get Services
```go
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  resp, err := api.Client.Scheduler.SchedulerGetServices(
    &scheduler.SchedulerGetServicesParams{
      Context: ctx,
    },
    api.WithAuth,
  )
  if err != nil {
    // handle err
  }


```

##### Search Availability
```go
  var (
    zip       = "92101"
    serviceID = "service uuid"
    rangeFrom = time.Now().Add(1 * time.Hour).Format(time.RFC3339)
    rangeTo   = time.Now().Add(3 * (24 * time.Hour)).Format(time.RFC3339)
    duration  = 15 // appointment duration length (15,30,45,60) 
  )

  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  resp, err := api.Client.Scheduler.SchedulerSearchAvailability(
    &scheduler.SchedulerSearchAvailabilityParams{
      Body: &models.SchedulerAvailabilitySearchRequest{
        RangeFrom: swag.String(rangeFrom),
        RangeTo:   swag.String(rangeTo),
        Zip:       swag.String(zip),
        ServiceID: serviceID,
        Duration:  swag.Int32(duration),
      },
      Context: ctx,
    },
    api.WithAuth,
  )
  if err != nil {
    // handle err
  }

```

##### Create an Appointment
```go
  var (
    apptID       = "<appointment ID from the availability search results>"
    refID        = "<your internal reference ID>"
    patientRefID = "<your internal patient reference ID>"
  )
  
  ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
  defer cancel()

  resp, err := api.Client.Scheduler.SchedulerCreateAppointment(&scheduler.SchedulerCreateAppointmentParams{
    Body: &models.SchedulingAppointment{
      ID:    apptID, // shilling reservation UUID
      RefID: refID,  // your reference ID
      Patient: &models.SchedulingPatient{
        ID:    "",            // shilling UUID (if not using ref id)
        RefID: patientRefID,  // your reference ID
      },
    },
    Context: ctx,
  },
    api.WithAuth,
  )

  if err != nil {
    // handle err
  }
````