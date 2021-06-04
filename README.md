# CodeChallenge
Code Challenge for Displaying Random Words

I was able to complete the coding part to display random words at specific intervals, and stop the process once it is completed the time interval.

I could only do the UI part until this point, I was trying to link the UI to server side but could not wrap up that part.

Thanks,
Sathya

****************************************************

<!DOCTYPE html>
<html lang="en">
<head>
    <title>Display Random Words</title>
</head>
<body>
  	<form id="form1" runat="server" name="RandomWords.go">
        <label>Enter Words (as comma separated): <textarea rows="3" cols="30" name="listOfWords" id="listOfWords"></textarea>
        </label>
        <input id="btnStart" type="submit" value="Start" OnClick="Start()">
        <input id="btnClear" type="reset" value="Clear">
    </form>
</body>
</html>
