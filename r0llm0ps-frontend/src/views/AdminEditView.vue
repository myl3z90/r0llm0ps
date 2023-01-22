<script setup>
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const machines = [
  {
    id: 1, 
    name: 'Machine One',
    status: 'Running'
  },
  {
    id: 2, 
    name: 'Machine Two',
    status: 'Sleeping'
  },
  {
    id: 3, 
    name: 'Machine Three',
    status: 'Sleeping'
  },
  {
    id: 4, 
    name: 'Machine Four',
    status: 'Sleeping'
  }
]

const route = useRoute()

let machineTmp = {}
if (route.params && route.params.id) {
    machines.forEach(m => {
        if (m.id == route.params.id) {
            machineTmp = m
        }
    })
}

const machine = ref(machineTmp)

</script>

<template>
    <router-link :to="{ name: 'admin' }">
        - Back
    </router-link>

    <div>
        <h2>Edit Machine</h2>
        <div>
            <form>
                <label>Name</label>
                <input type="text" name="name" :value="machine.name"/>
                <button type="submit">Update</button>
            </form>
        </div>
    </div>
</template>