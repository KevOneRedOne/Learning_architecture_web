// CommonJs
const fastify = require('fastify')({
    logger: true
});
const cors = require('@fastify/cors');

fastify.register(cors, {
    origin: '*'
})

fastify.register(require('@fastify/postgres'), {
    connectionString: `postgres://${process.env.POSTGRES_USER}:${process.env.POSTGRES_PASSWORD}@my-app-postgresql:5432/${process.env.DB_NAME}`
});

// Declare a route
fastify.get('/', function (request, reply) {
    reply.send({ hello: 'world' })
})

fastify.get('/api/articles', function (req, reply) {
    console.log(req.query);
    fastify.pg.query(
        'SELECT * FROM Articles',
        function onResult(err, result) {
            reply.send(err || result)    
        }
    )
})

fastify.get('/api/search', async function (req, reply) {
    console.log("Query : ",req.query);
    console.log("Query Param : ", req.query.title);
    try {
        const result = await fastify.pg.query(
            'SELECT * FROM articles WHERE title LIKE $1',
            [`%${req.query.title}%`]
        );
        reply.send(result.rows);
    } catch (err) {
        reply.send(err);
    }
})

// Run the server with docker - host 0.0.0.0
fastify.listen({ port: 8181, host: '0.0.0.0'}, function (err, address) {
    if (err) {
        fastify.log.error(err)
        process.exit(1)
    }
    fastify.log.info(`Server listening on ${address}`)
});
