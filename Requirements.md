### Minerva LMS

# Requirements

## Introduction

Minerva is a modern Learning Management System primarily for colleges and universities.

This document describes the complete requirements for Minerva version 1. Requirements are the features and functions of all applications and services in the system (functional) and the constraints under which they must operate (non-functional).

This is used by the development team to develop the system and by the designers to design the user interfaces.

Each requirement is given a unique _feature ID_ (fid) which is used to refer to the requirement.

## Groups

`fid:groups`

Groups are generalization of courses, faculties, departments, etc. They are used as a collection of resources and can have members.

### Group Types

`fid:groups.types`

A group type specifies:

- The name of the type
- An optional description
- An optional icon
- Whether groups are recurring or permanent

### Default Group Types

`fid:groups.types.defaults`

The default group types are:

- Course Offering
- Course
- Faculty
- Department
- University/College (which one depends on the config, see [Configuration](#configuration))

### Customizing Group Types

`fid:groups.types.customization`

The available types can be managed (edited, added, removed) by the admins.

### Group Hierarchy

`fid:groups.hierarchy`

The groups are organized in a hierarchy meaning that a department can have faculties, faculties can have courses, etc.

### Role Management

`fid:groups.role-management`

User roles are associated with groups, see [Roles](#roles). These roles can be managed (edited, added, removed) within groups.

### Group Members

`fid:groups.members`

Groups have members. Groups without members are archived, see [Archives](#archives).

Members are users and always have a certain role within the group, see [Roles](#roles).

### Joining Groups

`fid:groups.joining`

Users can join groups to become a member. Whether this is possible depends on the group settings, see [Group Settings](#group-settings).

### Requesting to Join Groups

`fid:groups.joining.request`

Depending on the group settings ([Group Settings](#group-settings)) users must be approved before they can join a group.

To be approved users must request to join the group. This request can be accepted or rejected by group members.

`fid:groups.joining.request.availability`

Whether users can even request to join a group depends on the group settings, see [Group Settings](#group-settings).

### Approving/Rejecting Requests

`fid:groups.joining.request.approval`

Which users can approve/reject requests depends on the group settings, see [Group Settings](#group-settings).

### Invitations

`fid:groups.invitations`

Group members can invite other users to join the group. Whether a user can do this depends on the group settings, see [Group Settings](#group-settings).

Depending on the group settings invited users must still also be approved before they can join a group.

`fid:groups.invitations.no-redundant-approval`

If a user is invited by a group member that is also allowed to approve requests, the user is automatically approved.

### Accepting/Rejecting Invitations

`fid:groups.invitations.acceptance`

Users can accept or reject the invitations to join a group.

### Leaving Groups

`fid:groups.leaving`

Users can leave groups to stop being a member. Whether this is possible depends on the group settings, see [Group Settings](#group-settings).

### Removing Members

`fid:groups.members.removing`

Group members can be removed by other group members. Which members can remove which remembers depends on the group settings, see [Group Settings](#group-settings).

### Group Settings

`fid:groups.settings`

Group settings specify:

- The available actions per role (including invitations, approvals/rejections, etc.)
- Which role in the parent group (or the parents parent group, etc.) is required to join or request to join the group
- ...

## Resources

`fid:resources`

Resources are content, i.e. documents, images, videos, etc. They exist within groups, see [Groups](#groups).

## Assignments

...

## Version Control

Minerva uses version control for all resources.

...

### Archiving

`fid:archives`

...

## Users

...

### Profile

...

### Roles

...

### Default Roles

`fid:users.roles.defaults`

The default roles are:

- Student
- Tutor
- Teacher
- Staff
- Admin

### Permissions

...

## Glossary

| Term | Definition                     |
| ---- | ------------------------------ |
| LMS  | Learning Management System     |
| User | Any person who uses the system |
