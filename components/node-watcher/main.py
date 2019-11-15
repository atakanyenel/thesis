from kubernetes import client, config,watch



def main():
    
    print("hello")
    config.load_incluster_config()
    api=client.CoreV1Api()

    nodes=api.list_node()
    for i in nodes.items:
        try:
            for t in i.spec.taints:
                if t.effect=="NoSchedule" and t.key=="node.kubernetes.io/unreachable":
                    api.delete_node(i.metadata.name)
                    print("Deleting node because it's not reachable")
        except TypeError:
            print("No taints")

    w=watch.Watch()
    for change in w.stream(api.list_node,timeout_seconds=0):
        node=change["object"]
        if change["type"]=="MODIFIED": # from ready to not ready
            try:
                for t in node.spec.taints:
                    if t.effect=="NoSchedule" and t.key=="node.kubernetes.io/unreachable":
                        api.delete_node(node.metadata.name)
                        print("Deleting node because it's not reachable")
            except TypeError:
                print("No taints")

if __name__=="__main__":
    main()
