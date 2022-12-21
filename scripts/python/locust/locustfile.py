import random

from locust import HttpUser, task


class Order(HttpUser):
    def on_start(self):
        self.client.post("/login", json={"username": "temp", "password": "temp"})

    @task
    def submit_order(self):
        num = str(random.randint(1, 5))
        self.client.post("/order/submit", json={"buildingNum": num})
