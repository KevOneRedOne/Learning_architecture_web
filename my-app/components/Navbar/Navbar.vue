<template>
    <div class="navbar">
        <NuxtLink to="/">
            <h1>My wonderfull app</h1>
        </NuxtLink>
        <nav>
            <ul>
                <li><NuxtLink to="/create">Créer un article</NuxtLink></li>
            </ul>
        </nav>
    </div>
    
</template>

<style>

    .navbar {
        background-color: #333;
        color: #fff;
        padding: 1rem;
        display: flex;
        justify-content: space-between;
    }

    nav {
        background-color: #333;
        color: #fff;
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 1rem;
    }

    ul {
        display: flex;
        list-style: none;
        margin: 0;
        padding: 0;
    }

    li {
        margin-left: 1rem;
    }

    a {
        color: #fff;
        text-decoration: none;
    }
</style>
