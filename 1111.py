import collections
def traverse(nodes, n):
    edges = collections.defaultdict(set)
    for u, v in nodes:
        edges[u].add(v)
        edges[v].add(u)
        
    visited = [False for _ in range(n)]
    
    stack = [] #Stack to store non visited node to explore later
    visited[0] = True
    stack.append(0)
    
    while stack:
        v = stack.pop()
        
        print(v)
        
        for u in edges[v]:
            if not visited[u]:
                visited[u] = True
                stack.append(u)
