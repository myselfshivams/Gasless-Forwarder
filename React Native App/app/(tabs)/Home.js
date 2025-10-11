const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 20,
    backgroundColor: "#121212", // darker background
  },
  title: {
    fontSize: 28,
    fontWeight: "bold",
    color: "#ff6f61", // accent title color
    marginBottom: 20,
    textAlign: "center",
    fontFamily: "Roboto, sans-serif",
  },
  walletContainer: {
    alignSelf: "center",
    width: "80%",
    backgroundColor: "rgba(255, 111, 97, 0.2)", // slightly transparent accent
    padding: 25,
    borderRadius: 100,
    marginBottom: 30,
    shadowColor: "#ff6f61",
    shadowOpacity: 0.3,
    shadowOffset: { width: 0, height: 4 },
    shadowRadius: 8,
  },
  walletText: {
    fontSize: 20,
    color: "#ffffff",
    textAlign: "center",
    fontWeight: "500",
  },
  formContainer: {
    flex: 1,
    justifyContent: "center",
  },
  tokenTypeContainer: {
    flexDirection: "row",
    justifyContent: "center",
    marginBottom: 30,
  },
  tokenButton: {
    paddingVertical: 16,
    paddingHorizontal: 30,
    backgroundColor: "rgba(255, 255, 255, 0.1)", // lighter buttons
    borderRadius: 30,
    marginHorizontal: 15,
    shadowColor: "#000",
    shadowOpacity: 0.1,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 5,
  },
  activeButton: {
    backgroundColor: "#ff6f61", // active button color
    elevation: 4,
  },
  buttonText: {
    color: "#ffffff",
    fontSize: 16,
    fontWeight: "bold",
    textAlign: "center",
  },
  label: {
    fontSize: 16,
    color: "#ff6f61", // label color changed to accent
    marginBottom: 8,
  },
  input: {
    backgroundColor: "rgba(255, 111, 97, 0.1)", // input background accent
    borderRadius: 50,
    marginBottom: 20,
    paddingHorizontal: 20,
    paddingVertical: 16,
    fontSize: 16,
    color: "#ffffff", // input text color
  },
  submitButton: {
    alignSelf: "center",
    width: "100%",
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: "#ff6f61", // main action color
    paddingVertical: 16,
    borderRadius: 50,
    marginTop: 10,
    shadowColor: "#ff6f61",
    shadowOpacity: 0.4,
    shadowOffset: { width: 0, height: 4 },
    shadowRadius: 8,
  },
  submitButtonText: {
    color: "#ffffff",
    fontSize: 18,
    fontWeight: "600",
  },
});
