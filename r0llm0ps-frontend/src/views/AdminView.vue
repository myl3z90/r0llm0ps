<script setup>
import { ref } from 'vue'

const machines = ref([
  {
    id: 1, 
    name: 'Machine One'
  },
  {
    id: 2, 
    name: 'Machine Two'
  },
  {
    id: 3, 
    name: 'Machine Three'
  },
  {
    id: 4, 
    name: 'Machine Four'
  }
])
</script>

<template>
    Hello Admin

    <div>
        <h2>New Machine</h2>
        <div>
            <form>
                <label>Name</label>
                <input type="text" name="name"/>
                <button type="submit">Create</button>
            </form>
        </div>
    </div>

    <div>
        <h2>Edit Machines</h2>
        <ul>
            <li v-for="machine in machines" :key="machine.id">
            {{ machine.name }} |
            <router-link :to="{ name: 'admin-edit', params: { id: machine.id }}">
                Edit | 
            </router-link>
            <a href="#">Delete</a>
            </li>
        </ul>
    </div>
</template>