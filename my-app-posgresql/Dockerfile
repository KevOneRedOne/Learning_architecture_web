FROM postgres:16.0-alpine3.18

# Copy the database initialization script:
# COPY .env /tmp/env.list
COPY ./scripts /docker-entrypoint-initdb.d/

# Set the environment variables:
ENV POSTGRES_DB ${POSTGRES_DB}
ENV POSTGRES_USER ${POSTGRES_USER}
ENV POSTGRES_PASSWORD ${POSTGRES_PASSWORD}

# Grant permissions for the initialization scripts:
RUN chmod a+r /docker-entrypoint-initdb.d/*.sql

# Expose the port:
EXPOSE 5432

# Run the database:
CMD ["docker-entrypoint.sh", "postgres"]


#Docker Build
#docker build -t m1_custom_postgres .

#Docker Run
#docker run -d --name m1_custom_postgres --env-file .env -p 5432:5432 -d m1_custom_postgres

# Start Postgres
# pg_ctl -D /var/lib/postgresql/data -l logfile start
