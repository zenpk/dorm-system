import random

from locust import HttpUser, task


class Order(HttpUser):
    @task
    def submit_order(self):
        num = random.randint(1, 5)
        self.client.post("/order/submit", json={"buildingNum": num})
