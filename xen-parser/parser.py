import json

with open("input.json") as f:

    data=json.load(f)

    for vm in data:
        try:
            print(vm["config"]["c_info"]["name"])
            print(vm["config"]["b_info"]["kernel"])
        except:
            pass