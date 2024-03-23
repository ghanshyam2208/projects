import app from './app-bootstrap';

const PORT = 3001;
app.listen(PORT, () => {
  console.log(`server is listening on ${PORT}`);
});
