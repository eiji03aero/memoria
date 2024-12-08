# Memoria
- An friendly application that tightens family bond.

# Development
```sh
# Start up backend
cd backend
./docker-compose.sh up

# Start emulator of mobile app
# Start ios device simulator first through launcher
cd mobile/memoria
flutter run --hot
```

# System
## memoria-api
- An api server which provides functionalities for memoria.

## memoria-client
- A client web application which is the interface of memoria for users.

## memoria-admin-api
- Api server to work with memoria administrative purpose
    - try to use RSC as much as possible to know the capacity of it

## memoria-admin-client
- Client web app for memoria admin
- Remarks
    - try to use RSC as much as possible to know the capacity of it
    - try look for admin oriented react package for a knowledge

# Reference
- [SPEC.md](docs/SPEC.md)
- [DESIGN.md](docs/DESIGN.md)
- [PLAN.md](docs/PLAN.md)
- [SELECTION.md](docs/SELECTION.md)
- [SERVICE.md](docs/SERVICE.md)
- Logo data: https://www.canva.com/design/DAGDb1mB68I/3ysozDgehKPU0oWjGwMzBA/edit
- Mobile app design: https://app.uizard.io/prototypes/ogO8Amq4PAiKxW7B8Oly
