# Use the official PostgreSQL 14 Alpine base image
FROM postgres:14-alpine

# Copy initialization scripts, if any
# COPY init.sql /docker-entrypoint-initdb.d/

# Expose the default PostgreSQL port
EXPOSE 5432

# Set the user to postgres
USER postgres
