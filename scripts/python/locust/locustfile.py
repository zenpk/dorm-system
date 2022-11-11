from locust import HttpUser, task


class Order(HttpUser):
    @task
    def hello_world(self):
        self.client.post("/order", json={"buildingNum": "1", "studentNum1": "1", "studentNum2": "2", "studentNum3": "3",
                                         "studentNum4": "4"})
