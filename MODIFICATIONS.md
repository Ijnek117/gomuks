# gomuks - Research Project Modifications

This repository is a modified fork of the original gomuks that can be found [here](https://github.com/gomuks/gomuks). The code here has been modified to serve as a component in the "Achieving User Pseudonymity in Federated End-to-End Encrypted Messaging Platforms" research project. For an overview of the project, see the [main project README](https://github.com/Ijnek117/...). 

---

## 1. Introduction

There was no public client implementation (that I could find) of [MSC4014: Pseudonymous Identities](https://github.com/matrix-org/matrix-spec-proposals/pull/4014), thus gomuks was chosen for the project as it is a minimal Matrix client that could be quickly modified to support rooms of this type and add support for encrypting invitee userIDs.

## 2. Summary of Changes

The following changes were made to this repository for the research project:

* Added support for inviting a user after a room was created in the Web client.
* Added support to encrypt a given userID with the certificate of their homeserver.

## 3. Rationale for Modifications

These changes were necessary to prevent a user's homeserver from tracking what remote users they are communicating with. By encrypting the localpart of the invitee's userID, a user's homeserver knows that they are inviting **some** user from homeserver B, but not which user it is.

## 4. Key Files Modified

This table provides a more detailed reference to the files and code that were changed or added.

| File Path | Change Description |
| :--- | :--- |
| `/go.mod` | Modified to use a local version of the maunium.net/go/mautrix dependency. |
| `/pkg/hicli/json-commands.go` | Modified the setMembership command for the "invite" action to first encrypt the userID using h.Client.EncryptUser and then send the invitation using h.Client.InviteUserWithResp with the encrypted ID. |
| `/web/src/api/types/mxtypes.ts` | Updated the RoomVersion TypeScript type to include "org.matrix.msc4014". |
| `/web/src/ui/menu/RoomMenu.tsx` | Added an "Invite a user" button and corresponding inviteUser function to the room context menu to allow invites after room creation. |