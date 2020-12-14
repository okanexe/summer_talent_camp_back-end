### Otsimo Summer Talent Camp Project

## Project Description
At Otsimo, we want to start a project called the Summer Talent Camp. At the beginning of the process, the 4th year university students from different departments will apply to the camp. After the selection progress, the interns will be given a task to work on during their internship. Among these interns, some will be hired as full-time team members at the end of the summer. To make the selection process simpler for the team, we would like to develop a platform that provides a way to manage candidates, their applications, and appointments relating to them easily and quickly. On this platform, the assigned team member will be able to quickly assess the candidates and see if there is anything that needs their attention. They will be able to accept or reject the application, find out if they have any appointments with the applicants, etc. As a developer, you will code the functions related to the database. We will use MongoDB as a database and GO as the programming language. All the details, rules are written below and an example DB dump (dumped with mongo 3.6 with this command mongodump --host localhost:27018 --archive=dump.gz --gzip --db Otsimo) is attached in description.

## Technical Details

## Objects
As you can check it from example DB we have two collections called Candidates and Assignees.
Candidates collection stores some required information of Candidates.
A candidate is defined as follows:

<i><b>_id string:</b></i> Unique hash that identifies candidate.<br/>
<i><b>first_name string:</b></i> First name of the candidate.<br/>
<i><b>last_name string:</b></i> Last name of the candidate.<br/>
<i><b>email string:</b></i> Contact email of candidate.<br/>
<i><b>department string:</b></i> Department that candidate applied.<br/>
Available values are<br/>
-Marketing<br/>
-Design<br/>
-Development<br/>
<i><b>university string:</b></i> University of the candidate.<br/>
<i><b>experience boolean:</b></i> Candidate has previous working experience or not.<br/>
<i><b>status string:</b></i> Status of the candidate.<br/>
Available values are<br/>
-Pending<br/>
-In Progress<br/>
-Denied<br/>
-Accepted<br/>
<i><b>meeting_count int:</b></i> The order of the next meeting. The maximum meeting count is 4.<br/>
<i><b>next_meeting DateTime:</b></i> Timestamp of the next meeting between the Otsimo team and the candidate.<br/>
<i><b>assignee string:</b></i> The id of the Otsimo team member who is responsible for this candidate.<br/>

Collection Assignees stores some required information relating to the Assignees. Assignee means a team member of Otsimo who is responsible for their own team's summer interns.

An assignee is defined as follows:

_id string: Unique hash that identifies an assignee.
name string: Name of the assignee.
department string: Department that assignee works in the Otsimo.
Available values are
-Marketing
-Design
-Development

## Function Signatures

The required storage functions are listed below with signatures.

CreateCandidate (candidate Candidate) (Candidate, error)
ReadCandidate (_id string) (Candidate, error)
DeleteCandidate (_id string) error
ArrangeMeeting (_id string, nextMeetingTime *time.Time) error
CompleteMeeting (_id string) error
DenyCandidate (_id string) error
AcceptCandidate(_id string) error
FindAssigneeIDByName (name string) string

## Requirements

Each candidate should have a unique identifier.

New candidate's Status should be Pending and meeting count should be 0. If meeting count is greater than 0 and smaller than 4, the Status should be In Progress.

New candidates should have an assignee who is working in the department that the candidate is applying to work.

Candidates cannot be accepted before the completion of 4 meetings.

If the next meeting is the last (4th) one, the assignee of the candidate should be changed to Zafer. He is the CEO of Otsimo.
