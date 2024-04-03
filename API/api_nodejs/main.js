const express = require('express');
require('dotenv').config();

const app = express();
const port = 8070;

app.get('/config', (req, res) => {
  const message = {
    log_level: process.env.LOG_LEVEL || 'DefaultLogLevel',
    grpc_port: process.env.GRPC_PORT || 'DefaultGRPCPort',
    environment: process.env.ENVIRONMENT || 'DefaultEnvironment',
    db_url: process.env.DB_URL || 'DefaultDBURL',
  };
  res.json(message);
});

app.get('/domain', (req, res) => {
  const host = req.get('host');
  console.log(`Host: ${host}`);

  const message = {
    endpoint: '/domain',
    domain: host || 'DefaultDomain',
    log_level: process.env.LOG_LEVEL || 'DefaultLogLevel',
  };
  res.json(message);
});

app.get('/load-capability', (req, res) => {
  const message = {
    cpu: process.env.CPU || 'DefaultCPU',
    memory: process.env.MEMORY || 'DefaultMemory',
  };
  res.json(message);
});

app.listen(port, () => {
  console.log(`Server is listening on port ${port}...`);
});

// created by Bexruz