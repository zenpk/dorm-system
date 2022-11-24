from locust import HttpUser, task


class Order(HttpUser):
    @task
    def hello_world(self):
        self.client.post("/test", json={})
