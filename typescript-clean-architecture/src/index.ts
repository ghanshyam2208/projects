import app from './app-bootstrap';

const PORT = 3000;
app.listen(PORT, () => {
  console.log(`server is listening on ${PORT}`);
});
