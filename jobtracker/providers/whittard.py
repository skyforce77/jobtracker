import requests
from pyquery import PyQuery
from .provider import Provider
from datetime import datetime

class Whittard(Provider):
    @staticmethod
    def get_jobs():
        jobs = []
        html_page = requests.get("https://careers.whittard.co.uk/contact/").text
        requester = PyQuery(html_page)
        all_jobs = requester(".job-cta a:first")
        for job in all_jobs:
            job_details = {}
            job_url = job.attrib["href"]
            html_page = requests.get(job_url).text
            requester = PyQuery(html_page)

            #TODO: Verify each value
            job_details["link"] = job_url
            job_details["title"] = requester(".position_title .jobs-row-input")[0].text
            job_details["description"] = requester(".position_description .jobs-row-input p:first")[0].text
            job_details["type"] = requester(".position_employment_type .jobs-row-input")[0].text
            job_details["city"] = requester(".position_job_location .jobs-row-input").remove('svg').text()
            job_details["date"] = datetime.strptime(
                requester(".type-date-posted .jobs-row-input").remove('svg').text(),
                "%B %d, %Y"
            )
            
            jobs.append(job_details)

        return jobs

