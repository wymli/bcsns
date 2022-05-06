// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

contract bcsns {
    event userMsgPersistedEvent(
        uint64 indexed send_uid,
        uint64 indexed recv_uid,
        int64 server_msg_id, 
        bytes message
    );
    event roomMsgPersistedEvent(
        uint64 indexed room_uid,
        uint64 send_uid,
        int64 server_msg_id,
        bytes message
    );
    event momentsPersistedEvent(
        uint64 indexed send_uid,
        int64 server_msg_id,
        bytes message
    );

    function PersistUserMessage(
        uint64 send_uid,
        uint64 recv_uid,
        int64 server_msg_id,
        bytes calldata message
    ) public {
        emit userMsgPersistedEvent(send_uid, recv_uid, server_msg_id, message);
    }

    function PersistRoomMessage(
        uint64 room_id,
        uint64 send_uid,
        int64 server_msg_id,
        bytes calldata message
    ) public {
        emit roomMsgPersistedEvent(room_id, send_uid, server_msg_id, message);
    }

    function PersistMoments(
        uint64 send_uid,
        int64 server_msg_id,
        bytes calldata message
    ) public {
        emit momentsPersistedEvent(send_uid, server_msg_id, message);
    }
}
