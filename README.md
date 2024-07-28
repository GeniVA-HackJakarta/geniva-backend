<br>
<div align="center">
    <div >
        <img height="150px" src="./img/logo.png" alt=""/>
    </div>
    <div>
            <h3><b>GeniVA</b></h3>
            <p><i><b>G</b>enerative AI for <b>En</b>hanced <b>I</b>nteractions and Experiences <b>V</b>irtual <b>A</b>ssistant
</i></p>
    </div>      
</div>
<br>
<h1 align="center">GeniVA Backend Service</h1>
The backbone of our solution: GeniVA's backend system powers the platform with robust and secure infrastructure, managing data seamlessly, facilitating real-time interactions, and ensuring smooth integration with AI. GeniVA streamlines the user experience, offers cost-effective travel options, provides personalized recommendations based on emotional states and situations, and encourages small, regular savings through collaboration with financial partners.

## 👨🏻‍💻 &nbsp;Technology Stack

<div align="center">

<a href="https://go.dev/">
<kbd>
<img src="https://miro.medium.com/v2/resize:fit:600/1*i2skbfmDsHayHhqPfwt6pA.png" height="60" />
</kbd>
</a>

<a href="https://gin-gonic.com/">
<kbd>
<img src="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" height="60" />
</kbd>
</a>

<a href="https://gorm.io/">
<kbd>
<img src="https://miro.medium.com/v2/resize:fit:1400/1*XBvxUxqycRC8B8KGCuzJVw.png" height="60" />
</kbd>
</a>

<a href="https://www.postgresql.org/">
<kbd>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Postgresql_elephant.svg/180px-Postgresql_elephant.svg.png" height="60" />
</kbd>
</a>

<a href="https://www.digitalocean.com/">
<kbd>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/f/ff/DigitalOcean_logo.svg/263px-DigitalOcean_logo.svg.png" height="60" />
</kbd>
</a>

<a href="https://www.docker.com/">
<kbd>
<img src="https://firebasestorage.googleapis.com/v0/b/upheld-acumen-420202.appspot.com/o/readme-assets%2Ficons%2FDocker.png?alt=media&token=3588896c-975f-496f-87d0-e7e1bce0d492" height="60" />
</kbd>
</a>

</div>
<div align="center">
<h4>Go | Gin | Gorm | PostgreSQL | Digital Ocean | Docker</h4>
</div>

## ⚙️ &nbsp;How to Run

1. Clone this repository from terminal using this following command
    ```bash
    $ git clone https://github.com/GeniVA-HackJakarta/geniva-backend.git
    ```
2. Create a `.env` file inside the repository directory using `.env.example` file as the template. You should add information about your own Google project to the `.env` file
3. Run the server using this following command, make sure you have Docker Desktop on your device.
    ```bash
    $ docker compose up
    ```
4. GeniVA backend server should be running. You can also check the server by opening http://localhost:8080

## 🔑 &nbsp;List of Endpoints

| Endpoint               | Method | Usage                                                  |
| ---------------------- | :----: | ------------------------------------------------------ |
| /api/auth/register     |  POST  | Users can register and create account on GeniVA        |
| /api/auth/login        |  POST  | Users can login to GeniVA using their existing account |
| /api/user              |  GET   | Users can get essential informations about users       |
| /api/user/:user_id/kyc |  PUT   | Users can update their KYC data                        |
| /api/history/:user_id  |  GET   | Users can retrieve their history data                  |
| /api/order/            |  POST  | Users can order Grab service                           |
| /api/ai/:user_id       |  POST  | Users can communicate with AI GeniVA                   |
| /api/price/calculate   |  POST  | Users can saving in their purchase to Superbank        |

## 👥 &nbsp;Contributors

| <a href="https://github.com/michaelsht"></a>                                                                                                        | <a href="https://github.com/NnA301023"></a>                                                                                            | <a href="https://github.com/TimothySubekti0322"></a>                                                                                        | <a href="https://github.com/AustinPardosi"></a>                                                                                                                  |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| <div align="center"><h3><b><a href="https://github.com/michaelsht">Michael Sihotang</a></b></h3><i><p>Bandung Institute of Technology</i></p></div> | <div align="center"><h3><b><a href="github.com/NnA301023">M. Alif Ramadhan</a></b></h3></a><p><i>Singaperbangsa Karawang</i></p></div> | <div align="center"><h3><b><a href="TimothySubekti0322">Timothy Subekti</a></b></h3></a><p><i>Bandung Institute of Technology</i></p></div> | <div align="center"><h3><b><a href="https://github.com/AustinPardosi">Austin Gabriel Pardosi</a></b></h3></a><p><i>Bandung Institute of Technology</i></p></div> |
