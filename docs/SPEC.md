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

## Idea features
- Timeline
    - maa twitter tekina yatsu
- Threads on medium
- Calendar view of album


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

---

### Create an album
#### Flow
- 1: Albums page
    - User can open create album page
    - Next
        - Open create album drawer to go 2.
- 2: Create album drawer
    - User submits form to create an album
    - Input
        - name
    - Next
        - Submit succeeds to 3.
- 3: Albums page
    - User will be feedbacked with the result

### View albums and view media uploaded
#### Flow
- 1: Albums page or All photo page
    - User can see the list of albums created
    - Next
        - Click album to open album page
- 2: Album page
    - User can see grid of media uploaded to album
    - Select one of the media
    - Next
        - Selects one medium succeeds to 3.
- 3: Medium page
    - User can view medium
    - User can move back and forth between media

### Upload media
#### Flow
- 1: Album page or All photo page
    - User selects album to open dedicated page
    - click on upload medium button
    - open file explorer and select media and submits
    - Input
        - any type of medium allowed
    - Next
        - Selects to submit to 2.
- 2: Album page
    - User will be feedbacked with the result
#### Remark
- Use signed url to upload so that api server won't have to touch the media files
- System flow:
    - client gets media

### Delete media
#### Flow
- 1: Album page or all photo page
    - User selects album to open dedicated page
    - Start selection
    - User will pushes menu bottom right and select delete media
    - Input
        - media ids
    - Next
        - confirm request succeeds to to 2.
- 2: Album page
    - User will be feedbacked with the result

### Add media to albums
#### Flow 1
- 1: Album page or All photo page
    - User selects photos
    - User use menu item on bottom right
    - Next
        - Click menu item of add to album to go 2.
- 2: Add to album drawer
    - User selects albums to add to
- 3: Album page
    - User will be feedbacked with result

### Remove media from albums
#### Flow 1
- 1: Album page
    - User selects photos
    - User use menu item on bottom right
    - Next
        - Click menu item of unlink to album to go 2.
- 2: Unlink to album drawer
    - Next
        - Confirm the action succeeds to 3.
- 3: Album page
    - User will be feedbacked with result

---

### Get timeline for user space
- General
    - User can post
    - Activities are posted as well
        - uploaded photo
        - user joined

---

### Create another user space with existing user
#### Flow
- 1: Configuration page
    - User can open create another user space page
    - Next
        - Click menu to go 2.
- 2: Create user space page
    - User can fill out form to create another user space page

---

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

---

### Create an album
- POST /api/auth/albums
- Input:
    - name: string

### Get list of albums
- GET /api/auth/albums

### Get album detail
- GET /api/auth/albums/:id
- Input:
    - id: string

### Request media upload urls
- POST /api/auth/media/request-upload-urls
- Input:
    - file_names: string[]
    - album_id?: string

### Get list of media
- GET /api/auth/media
- Input:
    - album_id?: string

### Get page of media
- GET /api/auth/media/get-page
- Input
    - album_id?: string
    - medium_id: string
- Output
    - current_page: int
    - total_page: int

### Confirm uploads
- POST /api/auth/media/confirm-uploads
- General
    - Confirm the uploads to api server so that the server can do wrap work like:
        - create thumbnails
- Input:
    - medium_ids: []string

### Delete media
- DELETE /api/auth/media/:id

### Delete media
- DELETE /api/auth/albums/:id

### Add media to albums
- POST /api/auth/albums/:id/add-media
- Input
    - medium_ids: []string
    - album_ids: []string

### Remove media from album
- POST /api/auth/albums/:id/remove-media
- Input
    - medium_ids: []string
    - album_ids: []string

---

# Reminding notes
## Workmail
- It is created in oregon region
