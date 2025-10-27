---
title: "Web API and Model/DTO ... | BookStack - RNT"
source: "https://bookstack.rnt.lscm/books/general/page/web-api-and-modeldto-naming-guide"
author:
published:
created: 2025-10-27
description: "General API Path/api/<Entity>/<API Name>The entity can be a table name or a group name, and it m..."
tags:
  - "clippings"
---

### General API Path

/api/< Entity >/< API Name >

The entity can be a table name or a group name, and it may contain child entities.

<table><colgroup><col> <col></colgroup><tbody><tr><td colspan="2">Example</td></tr><tr><td><strong>Table Name (Most of the cases)</strong></td><td><strong>Group Name</strong></td></tr><tr><td><ul><li>/api/ <span>user</span> / <span>create</span></li><li>/api/ <span>user</span> / <span>create <span>Role</span></span></li><li>/api/ <span>sensor</span> / <span>update</span></li></ul></td><td><ul><li>/api/ <span>auth</span> / <span>login</span></li><li>/api/ <span>data</span> / <span>qoeSensorPhotoFile</span></li></ul></td></tr></tbody></table>

By default, the API should target an entity, so "/api/ user / create " means create a user. On the other hand, it can target a child entity. For example, "/api/ user / create Role " means " role " is a child entity of user.

Do not share the DTO request and response objects

### Server & Client Model/DTO definition

|            | Request DTO                   | Response DTO     | Internal DTO |
| ---------- | ----------------------------- | ---------------- | ------------ |
| **Server** | Master copy                   | Master copy      | Master copy  |
| **Client** | **N/A (Or copy from server)** | Copy from Server | Master copy  |

### API to retrieve general information

| Abbreviation | Return type        | Empty result    | Description                                                   |
| ------------ | ------------------ | --------------- | ------------------------------------------------------------- |
| **`qry`**    | List of objects    | Empty List      | " Q ue ry ", a general query, and it must return a list       |
| **`qon`**    | One object or null | null            | " Q uery O ne N ull safe", query one object and allow null    |
| **`qoe`**    | One object         | Throw exception | " Q uery O ne E xact", query one object and do not allow null |

##### Paths:

- /api/ <Entity> / <qry/qon/qoe>
- /api/ <Entity> / <qry/qon/qoe> <Description, By XXX, With XXX>
- /api/ <Entity> / <qry/qon/qoe> **Of** <Child entity> <Description, By XXX, With XXX>

##### Query Example

- /api/ user / qry ByLocation
- /api/ data / qry **Of** SensorBattery ByLocation

##### Query One Null Safe Example

- /api/ locations / qonMapFile

##### Query One Exact Example

- /api/ user / qoe
- /api/ user / qoePg
- /api/ user / qoeDetailPg
- /api/ user / qoeEditPg
- /api/ user / qoeCreatePg
- /api/ dashboard / qoePg
- /api/ dashboard / qoeSensorPgs
- /api/ dashboard / qoeAlertPgs

#### Request DTO

- EmptyReq
- IdReq
- < Entity > <Inf/Pg/Pgs> Req
- < Entity > <Description, By XXX, With XXX> <Inf/Pg/Pgs> Req
- < Child entity > <Inf/Pg/Pgs> Req
- < Child entity > <Description, By XXX, With XXX> <Inf/Pg/Pgs> Req

#### Response DTO

| Type                      | Request     | Abbreviation  | Response DTO Name                            | Description  |
| ------------------------- | ----------- | ------------- | -------------------------------------------- | ------------ |
| General query             | qon/qoe/qry | Inf (default) | <Entity> <Description, By XXX, With XXX> Inf | Query Item   |
| View Model - Page         | qoe         | Pg            | <Entity> <Description, By XXX, With XXX> Pg  | Page         |
| View Model - Page session | qon/qoe     | Pgs           | <Entity> <Description, By XXX, With XXX> Pgs | Page Session |

#### Example for a dashboard page

[![image.png](https://bookstack.rnt.lscm/uploads/images/gallery/2025-04/scaled-1680-/h7Cimage.png)](https://bookstack.rnt.lscm/uploads/images/gallery/2025-04/h7Cimage.png)

Case 1: Single API call

| API                  | Request        | Response    |
| -------------------- | -------------- | ----------- |
| /api/dashboard/qoePg | DashboardPgReq | DashboardPg |

Case 2: Multi session API call

| API                             | Request                 | Response             |
| ------------------------------- | ----------------------- | -------------------- |
| /api/dashboard/qoeOfMapPgs      | DashboardMapPgsReq      | DashboardMapPgs      |
| /api/dashboard/qoeOfSensorPgs   | DashboardSensorPgsReq   | DashboardSensorPgs   |
| /api/dashboard/qoeOfAlertPgs    | DashboardAlertPgsReq    | DashboardAlertPgs    |
| /api/dashboard/qoeOfLocationPgs | DashboardLocationPgsReq | DashboardLocationPgs |

Case 3: Multi info API call

| API                          | Request                 | Response                     |
| ---------------------------- | ----------------------- | ---------------------------- |
| /api/dashboard/qryOfMap      | DashboardMapQryReq      | List of DashboardMapInf      |
| /api/dashboard/qryOfSensor   | DashboardSensorQryReq   | List of DashboardSensorInf   |
| /api/dashboard/qryOfAlert    | DashboardAlertQryReq    | List of DashboardAlertInf    |
| /api/dashboard/qryOfLocation | DashboardLocationQryReq | List of DashboardLocationInf |

Case 4: Multi info API call (Different entity)

| API                              | Request                    | Response                           |
| -------------------------------- | -------------------------- | ---------------------------------- |
| /api/location/qryForDashboardMap | LocationForDashboardMapReq | List of LocationForDashboardMapInf |
| /api/sensor/qryForDashboard      | SensorForDashboardQryReq   | List of SensorForDashboardInf      |
| /api/alert/qryForDashboard       | AlertForDashboardQryReq    | List of AlertForDashboardInf       |
| /api/location/qryForDashboard    | LocationForDashboardQryReq | List of LocationForDashboardInf    |

### API to retrieve Combo Box options

**Abbreviation:**`**cmb**`

##### Paths:

- Simple Entity /api/< Entity >/ cmb
  - /api/ sensorType / cmb
- Entity with Description /api/< Entity >/ cmb < Description e.g. By XXX, With XXX, etc >
  - /api/ sensorType / cmb ByCategoryId
- Child Entity with Description /api/< Entity >/ cmb **Of** <Child entity> <Description e.g. By XXX, With XXX, etc>
  - /api/ sensors / cmb **Of** SensorType CodeByCategoryId

##### Request DTO:

- EmptyReq
- IdReq
- < Entity > Cmb Req
- < Entity > Cmb2 Req
- < Entity > <Description, By XXX, With XXX> Cmb Req
- < Child entity > Cmb Req
- < Child entity > <Description, By XXX, With XXX> Cmb Req

##### Response DTO (List):

- < Entity > Cmb
- < Entity > <Description, By XXX, With XXX> Cmb
- < Child entity > Cmb
- < Child entity > <Description, By XXX, With XXX> Cmb

### API to retrieve information for Generic List (Our Custom component)

A specific API for the Generic List component to query the data. It should support filtering, sorting, and paging functions.

**Abbreviation:**`**lst**`

##### Paths:

- Simple Entity /api/< Entity >/ lst
  - /api/ sensorType / lst
- Entity with Description /api/< Entity >/ lst <Description>
  - /api/ sensorType / lst ForDashboard
  - /api/ sensorType / lst 2
  - /api/ sensorType / lst 3
- Child Entity /api/< Entity >/ lst **Of** <Child entity>
  - /api/ sensors / lst **Of** SensorType
- Child Entity with Description /api/< Entity >/ lst **Of** <Child entity> <Description>
  - /api/ sensors / lst **Of** SensorTypeForDashboard

##### Request DTO:

- ListReq

##### Response DTO (List of Information inf):

- - < Entity > **`Inf`**
  - < Entity > <Description> **`Inf`**
  - < Child entity > **`Inf`**
  - < Child entity > <Description> **`Inf`**

### API to perform an operation

Create, Update, Delete, Add, Insert, Generate, etc...

##### Paths:

- Simple Entity /api/< Entity >/ <operation>
  - /api/ sensorType / create
  - /api/ sensorType / update
  - /api/ sensorType / delete
  - /api/ sensorType / save
  - /api/ sensorType / generate
- Entity with Description /api/< Entity >/ <operation> < Description e.g. By XXX, With XXX, etc >
  - /api/ sensorType / create ByCode
- Child Entity with Description /api/< Entity >/ <operation> **Of** <Child entity> <Description e.g. By XXX, With XXX, etc>
  - /api/ sensors / create **Of** SensorType ByCode

##### Request DTO:

- EmptyReq
- IdReq
- < Entity > <operation> Req
  - UserCreateReq
- < Entity > <Description, By XXX, With XXX> <operation> Req
  - UserDeleteByUsernameReq
- < Child entity > <operation> Req
  - UserRoleAddReq
- < Child entity > <Description, By XXX, With XXX> <operation> Req

##### Response DTO:

It should not have any response DTO normally

- < Entity > <operation> Res
  - ReportGenerateRes
- < Entity > <Description, By XXX, With XXX> <operation> Res
- < Child entity > <operation> Res
- < Child entity > <Description, By XXX, With XXX> <operation> Res

### API to perform multiple operations

Example:

- UserCreateAndQoe
  - Return UserInf
- UserUpdateAndQoe
