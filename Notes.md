# TODO

-   [ ] Landing page

    -   [ ] Testimonials

-   [ ] Find Location (Don't use Google Maps API yet)
-   [ ] Services
-   [ ] Contact

    -   [ ] Send Email (anon)

-   [ ] Admin Dashboard

    -   [ ] Locations
    -   [ ] Data Visualizations
    -   [ ] Employees
    -   [ ] Dues and ...

-   [ ] User Dashboard

    -   [ ] First Name
    -   [ ] Last Name
    -   [ ] License plate no.
    -   [ ] Email address
    -   [ ] Car color
    -   [ ] Service bought (with date)
    -   [ ] Subscription program
    -   [ ] Membership/Subscription status

-   [ ] Payment
-   [ ] Implement `Auth` middleware

# Data Modelling

## Brainstorming

```
{
    location: {
        loc_stats: {
            tot_revenue: number,
            tot_investment: number
        },
        due: number,
        employees: [<<elmployee>>]
        vehicles: [<<com_vehicles>>]
    },
    daily_stats: {
        loc: <<location>>,
        revenue: number,
        invested: number
        customer_count: number,
    },
    customer: {
        car_info: {},
        customer_info: {}
        from_loc: <<location>>
    },
    employee: {
        name: string,
        salary: number,
        job: string,
        date_hired: date,
        pt_ft: bool,
        payment_due: number
        assigned_loc: <<location>>
    },
    gear: {
        cost: number,
        quantity: number
    },
    com_vehicles: {
        vehicle_id: string,
        license_plate: string,
        assigned_to: <<employee>>
    },
    services: {
        name: string,
        price: number,
        frequency: number
    }
}
```

# Endpoints

-   PUT `/check-in`
-   PUT `/check-out`
-   POST `/make-appointment`
-   GET `/see-unconfirmed-appointments`
-   GET `/see-confirmed-appointments`
-   PUT `/assign-employee-to-appointment`
-   PUT `/confirm-appointment`
-   PUT `/assign-employee-to-date`
-   PUT `/cant-make-date`
-   PUT `/volunteer`
-   POST `/apply-for-job`
