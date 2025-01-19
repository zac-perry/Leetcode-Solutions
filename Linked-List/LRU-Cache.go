// 146. LRU Cache
// Approach here was to use a hashmap for quick key lookup (O(1)
// I then used a doubly linked list to keep track of the cache entries. Map would allow me to find these entries super fast. When removing / inserting, was just a matter of swapping some nodes. 
// Lastly, I also always evicted the last entry in the doubly linked list, which indicated that this entry was the LRU since new entires + recently gotten entries were moved to the front
type LRUCache struct {
    entries map[int]*EntryNode
    sentinel *EntryNode
    capacity int
}

type EntryNode struct {
    key int
    val int
    prev *EntryNode
    next *EntryNode
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        make(map[int]*EntryNode, 0),
        &EntryNode{prev: nil, next: nil},
        capacity,
    }
}

func (this *LRUCache) PushFront(node *EntryNode) {
    if this.sentinel.next != nil {
        this.sentinel.next.prev = node
        node.next = this.sentinel.next
        node.prev = this.sentinel
        this.sentinel.next = node
        return
    }

    this.sentinel.next = node
    node.prev = this.sentinel
    node.next = this.sentinel
    this.sentinel.prev = node

    return
}

func (this *LRUCache) MoveToFront(node *EntryNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
    node.prev = this.sentinel
    node.next = this.sentinel.next
    this.sentinel.next.prev = node
    this.sentinel.next= node

    return
}

func (this *LRUCache) PopBack() int {
    backnode := this.sentinel.prev
    this.sentinel.prev = backnode.prev
    backnode.prev.next = this.sentinel
    backnode.next = nil
    backnode.prev = nil
    
    return backnode.key
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.entries[key]; ok {
        // if the key exists.
        // 2 things: 
            // 1., if its prev is not the sentinel, move it to the front.
            // return the val
        if node.prev != this.sentinel {
            this.MoveToFront(node)
        }
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int)  {
    // Call get, see if it exists. 
        // if so, update the value, move to the front of the linked list
    val := this.Get(key)
    if val != -1 {
        if node, ok := this.entries[key]; ok {
            node.val = value
            // if not already at the front of the cache
            if node.prev != this.sentinel {
                this.MoveToFront(node)
            }
           return 
        }
    }
    // If it does NOT exist.
    // if full, remove node from the back. sentinel->prev.
        // Also remove from map
    // Then insert
    // increment capacity.
    if len(this.entries) == this.capacity {
        // pop off the back
        // also remove from the map
        removedKey := this.PopBack()
        delete(this.entries, removedKey)
    } 

    // insert into the front, add to map, increment capacity
    node := &EntryNode{
        key,
        value,
        nil,
        nil,
    }
    this.PushFront(node)
    this.entries[key] = node

    return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 *
