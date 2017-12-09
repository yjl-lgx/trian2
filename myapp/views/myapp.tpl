<!DOCTYPE html>

<html>
<body>
  <header>
    <h1 class="logo">Myapp</h1>
  Products:{{.Products}} </br>
  Edition:{{.Edition}} </br>
   </br>
    </br>
    <div class="description">
      Myapp is a my feist one web. </br>
      my  ip: "{{.My_app}}" </br>
      UserAgent : "{{.User}}" </br> 
      My_hostname: "{{.My_hostname}}"</br>
      </br>
        </br>
          </br>
            </br>
              </br>
                </br>

      {{.username}} </br>
      {{.introduce}} </br>
        </br>
          </br>
            </br>
              </br>
    </div>
  </header>
  <footer>
    <div class="author">
      Official website:
      <a href="http://{{.err}}">{{.err}}</a> /
      Contact me:
      <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
    </div>
  </footer>
  <div class="backdrop"></div>


</body>
</html>