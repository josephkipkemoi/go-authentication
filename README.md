## GO-AUTH API

## Setup

## Dependencies 
<ul>
    <li>go 1.21.6</li>
</ul>

## Getting Started
Setting up the project in development mode:
<ul>
    <li>Ensure go is installed by running:</li>
    <p>go version</p>
    <p>Expected output: go version go1.21.6 (*OS)</p>
</ul>

## Running the Application
<p>Inside the root folder, run the following commands in your terminal</p>
<ul>
    <li>go run main.go</li>
</ul>

## Running the Tests
<ul>
    <li>go test ./tests</li>
 </ul>

 ## API
 <table>
    <h2>USER REGISTRATION AND AUTHENTICATION</h2>
    <tr>
        <td>METHOD</td>
        <td>END POINT</td>
        <td>PARAMS</td>
        <td>PARAM REQUIREMENTS</td>
    </tr>
    <tr>
        <td>POST</td>
        <td>/api/register</td>
        <td>
            json{"phone_number": 700545727, "password": "123456", "confirm_password": "123456"}
        </td>
        <td>
            phone_number -> 9 digits
            password -> Minimum 6 characters
        </td>
    </tr>
 </table>