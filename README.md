sent rpc message from user__1 to user user__2
```bash
{"method":40,"request_id":"user__2","body":{"data":"test"}}
```

after message was sent, recipient will receiver 2 events:
- message 
- message sent
```bash
{"type":20,"response_to":"user__2","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b","data":"test"}}
{"type":21,"response_to":"user__2","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b"}}
```

sent rpc message delivered
received event that message delivered
```bash
{"method":41,"request_id":"user__1","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b"}}
{"type":22,"response_to":"user__2","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b"}}
```

sent rpc message read
received event that message read
```bash
{"method":42,"request_id":"user__1","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b"}}
{"type":23,"response_to":"user__2","body":{"message_id":"b6c9aa8b-e864-499f-8c85-e4c8016cc13b"}}
```

sent rpc message typing start
received event that user typing start
```bash
{"method":60,"request_id":"user__1"}
{"type":40,"response_to":"user__2"}
```

sent rpc message typing end
received event that user typing end
```bash
{"method":61,"request_id":"user__1"}
{"type":41,"response_to":"user__2"}
```