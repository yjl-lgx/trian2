<html>
<body style="background-color:black">
<title>登陆界面</title>

<form method="post">
<center>
<h1 style="background-color:red">login</h1>
</br>
</br>
</br>
  Products:{{.Products}} </br>
  Edition:{{.Edition}} </br>
   </br>
    </br>
<div class="login" style="position:absolute; top:50%;left:50%;margin-top:-100px;margin-left:-200px;background:#69F; width:400px; height: 200px; border-radius:10px;">
      <div style=" margin-left:100px;margin-top:50px">
<table>           
<td><font color="black"><pre>User name</pre></font></td>
<td><input name="username" placeholder="Please enter number" maxlength="8" /></td>
</table> 
      </div>
    <div style="margin-left:100px;margin-top:20px">
<table>           
<td><font color="black"><pre>Password </pre></font></td>
<td><input type="password" name="password"  placeholder="Password" maxlength="8"/> </td>
</table> 
    </div>
      <div style=" text-align:center;margin-top:20px">
        <input type="submit" value="login" style=" padding:5px 15px; background:#F9C; border-radius:3px; border:#F0C 1px solid;"/>
         
      </div>
</div>
</center>
</form>
</body>
</html>
