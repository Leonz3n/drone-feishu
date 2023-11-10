# drone-feishu
Push message to Feishu in Drone Pipeline.

```yaml
steps:
  - name: Notice to Feishu
    image: leonz3n/drone-feishu
    settings:
      webhook: xxxxxxx
      secret: xxxxx
```