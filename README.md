# Memoria
- An friendly application that tightens family bond.

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

# Features
## Main features
- Dashboard
    - Gives overview of activities, user space
- Albums
    - can view, upload and organize media
- Slides
    - can create and share your memory in slides
- Setting
    - do your thing

## TBD
- Timeline
    - maa twitter tekina yatsu

## Usecases
### Signup
#### Flow
- 1: Sign up page
    - User fills form to create account and user space
    - Inputs
        - user name
        - user space name
        - email
        - password
    - Next
        - Submit succeeds to 2.
    - Remarks
        - Cannot use existing email
- 2: Signup email guide page
    - User will be guided to confirm email
    - Next
        - Opening email account succeeds to 3.
- 3: Check signup invitation email
    - User checks invitation email on their mailbox
    - User clicks on confirmation url
    - Next
        - Clicking confirmation url succeeds to 4.
- 4: Signup thanks page
    - User will be thanked for signup
    - The page also shows top page link

#### Remark
- If user wants to open new user space with existing user, need to do so on dedicated feature
    - not gonna create it for now
- User not confirmed email cannot acccess auth pages

#### Test
- 1: Signup success
    - do: execute signup form
    - check: redirected to signup email confirmation guide page
    - check: db user created with:
        - values in form
        - invited account status
    - check: open email
    - do: open confirm link
    - check: redirected to signup thanks page
    - check: db user updated with
        - confirmed account status
- 2: Protect auth pages with user account status
    - do: execute signup form
    - do: open top page
    - check: redirected to lp

### Login
#### Flow
- 1: Open login page
    - User opens login page from lp
    - Next
        - Succeeds to 2.
- 2: Submits login form
    - User fills out the form and submits to login
    - Next
        - Submit succeeds to 3.
- 3: User gets redirected to dashboard page
    - User gets redirected

#### Test
- 1: Login succeeds
    - do: open login page
    - do: fill out and submit login form
    - check: gets redirected to dashboard page
- 2: Validation error shown 1
    - do: open login page
    - do: fill out with unexisting email
    - do: submit the form
    - check: validation message related to the above shown
- 3: Validation error shown 2
    - do: open login page
    - do: fill out with existing email
    - do: fill out incorrect password
    - do: submit the form
    - check: validation message related to the above shown

### Logout
#### Flow
- 1: Opens account page
    - User opens account page to use logout link button
    - Next
        - Use logout link button succeeds to 2.
- 2: User logouts
    - User logouts and gets redirected to login page

#### Remark
- Let's execute the logic on frontend

#### Test
- 1: Logout succeeds
    - do: open account page and use logout button
    - check: gets redirected to login page
    - check: try to open auth page and confirm getting redirected to login page

### Invite users
#### Flow
- 1: Invite user page
    - User fills form to invite other users
    - Inputs
        - email
    - Next
        - Submit succeeds to 2.
    - Remarks
        - Cannot invite email that exists in the same user space
- 2: Dashboard page
    - User will be feedbacked with the result
- 3: User receives email
    - User receives email to sign up
    - Next
        - Open url to go 4.
    - Remark
            - Has token to differentiate invitation
- 4: User sign up by invitation
    - User fills form to sign up
    - Inputs
        - email (disabled just to show)
        - name
    - Next
        - Submit succeeds to 5.
- 5: Top page
    - User will be feedbacked with the result

### Create an album
#### Flow
- 1: Albums page
    - User can open create album page
    - Next
        - Open create album page to go 2.
- 2: Create album page
    - User submits form to create an album
    - Input
        - name
    - Next
        - Submit succeeds to 3.
- 3: Opens created album page
    - User will be feedbacked with the result

### View albums and view media uploaded
#### Flow
- 1: Albums page
    - User can see the list of albums created
    - Next
        - Click album to open album page
- 2: Album page
    - User can see grid of media uploaded to album

### Upload media
#### Flow
- 1: Album page
    - User selects album to open dedicated page
    - Click add media button, which allows selecting media
    - Input
        - any type of media allowed
    - Next
        - Selects to submit to 2.
- 2: Album page
    - User will be feedbacked with the result

#### Remark
- Use signed url to upload so that api server won't have to touch the media files

### Link existing media to albums
#### Flow
- 1: Media detail page
    - User opens media detail page
    - There is a menu to trigger linking media to album
    - Next
        - Selects album to link to go 2.
- 2: Medi adetail page
    - User will be feedbacked with the result

### Comment thread on media
#### Flow
- 1: Media detail page
    - User opens media detail page to open thread drawer
    - Next
        - Opens thread drawer to go 2.
- 2: Media detail page
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
- 1: Calendar page
    - User can view which date has which media taken
    - Next
        - Click on date to open date detail page
- 2: Date detail page
    - User can view grid of media for the given date
    - User can also:
        - Go back to calendar page
        - Go back and forth date

#### Remark
- The date to link should be the date it was taken, not the date uploaded

### Create another user space with existing user
#### Flow
- 1: Configuration page
    - User can open create another user space page
    - Next
        - Click menu to go 2.
- 2: Create user space page
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

### Invite user
- POST /api/auth/invte-user
- Input:
    - email: string
- Remark:
    - email has to be unique among all users

### Confirm invite user
- GET /api/public/invite-user-confirm
- Input:
    - invitation_id: string
    - name: string
    - password string

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
- Output:
    - token: string
- Remarks
    - email has to be unique among all users
    - jwt will be attached in response header

### Confirm signup
- GET /api/public/signup-confirm
- General:
    - Confirms signup
    - Verify user account with email
- Input:
    - id: string
- Redirects:
    - client/signup-thanks

### Login
- POST /api/public/login
- General:
    - This is to login for user
- Input:
    - email: string
    - password: string

# Reminding notes
## Workmail
- It is created in oregon region

# Issues
## service or dao?
- especially query logics
- they are kind of scattered among these two
- should have solid rules or reorganization

# Reference
- Logo data: https://www.canva.com/design/DAGDb1mB68I/3ysozDgehKPU0oWjGwMzBA/edit
