# Memoria
- An friendly application that tightens family bond.

# Features
## Usecases
### Create account
#### Flow
- 1. Sign up page
    - User fills form to create account and user space
        - Inputs
            - user space name
            - email
            - name
            - password
        - Next
            - Submit succeeds to 2.
        - Remarks
            - Cannot use existing email
- 2. Top page
    - User will be presented with top page

#### Remark
- If user wants to open new user space with existing user, need to do so on dedicated feature
    - not gonna create it for now

### Invite users
#### Flow
- 1. Invite user page
    - User fills form to invite other users
        - Inputs
            - email
        - Next
            - Submit succeeds to 2.
        - Remarks
            - Cannot invite email that exists in the same user space
- 2. Top page
    - User will be feedbacked with the result
- 3. User receives email
    - User receives email to sign up
        - Next
            - Open url to go 4.
        - Remark
            - Has token to differentiate invitation
- 4. User sign up by invitation
    - User fills form to sign up
        - Inputs
            - email (disabled just to show)
            - name
        - Next
            - Submit succeeds to 5.
- 5. Top page
    - User will be feedbacked with the result

### Create an album
#### Flow
- 1. Albums page
    - User can open create album page
    - Next
        - Open create album page to go 2.
- 2. Create album page
    - User submits form to create an album
    - Input
        - name
    - Next
        - Submit succeeds to 3.
- 3. Opens created album page
    - User will be feedbacked with the result

### View albums and view media uploaded
#### Flow
- 1. Albums page
    - User can see the list of albums created
    - Next
        - Click album to open album page
- 2. Album page
    - User can see grid of media uploaded to album

### Upload media
#### Flow
- 1. Album page
    - User selects album to open dedicated page
    - Click add media button, which allows selecting media
    - Input
        - any type of media allowed
    - Next
        - Selects to submit to 2.
- 2. Album page
    - User will be feedbacked with the result

#### Remark
- Use signed url to upload so that api server won't have to touch the media files

### Link existing media to albums
#### Flow
- 1. Media detail page
    - User opens media detail page
    - There is a menu to trigger linking media to album
    - Next
        - Selects album to link to go 2.
- 2. Medi adetail page
    - User will be feedbacked with the result

### Comment thread on media
#### Flow
- 1. Media detail page
    - User opens media detail page to open thread drawer
    - Next
        - Opens thread drawer to go 2.
- 2. Media detail page
    - User can do following stuff on thread drawer
        - send text message
        - see messages sent to thread
            - name
            - datetime
            - message content

#### Remark
- Place refresh button to load latest messages

### View media in calendar
#### Flow
- 1. Calendar page
    - User can view which date has which media taken
    - Next
        - Click on date to open date detail page
- 2. Date detail page
    - User can view grid of media for the given date
    - User can also:
        - Go back to calendar page
        - Go back and forth date

#### Remark
- The date to link should be the date it was taken, not the date uploaded

### Create another user space with existing user
#### Flow
- 1. Configuration page
    - User can open create another user space page
    - Next
        - Click menu to go 2.
- 2. Create user space page
    - User can fill out form to create another user space page

# Design
## Tech stack
- Frontend
    - Next.js
    - React-spectrum by adobe
    - react-i18next
    - PWA
        - installable
        - offline mode sounds good
- Backend
    - Golang
- Middleware
    - Postgres
- Infra
    - aws

## API
### Get app data
- GET /api/auth/app-data
- General
    - Provides data globally used on app
- Output:
    - user
        - id
        - name
    - user_space
        - id
        - name

### Signup
- POST /api/public/signup
- General:
    - Creates user
    - Creates user space
    - Link user and user space created
- Input:
    - name: string
    - email: string
    - password: string
    - user_space_name: string
- Remarks
    - email has to be unique among all users
    - jwt will be attached in response header

# Reference
- Logo data: https://www.canva.com/design/DAGDb1mB68I/3ysozDgehKPU0oWjGwMzBA/edit
