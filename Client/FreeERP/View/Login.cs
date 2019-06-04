using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace FreeERP.View
{
    public partial class Login : Form
    {

        MainPage p_MainPage;
        

        public Login(MainPage mp)
        {
            InitializeComponent();
            this.p_MainPage = mp;
        }

        private void Login_Load(object sender, EventArgs e)
        {

        }

        private void showMainPage()
        {
            this.p_MainPage.Visible = true;
        }

        // 登录 
        private void LoginClick(object sender, EventArgs e)
        {
            if(username.Text == "admin" && password.Text == "123")
            {
                this.Visible = false;
                this.showMainPage();
            } else
            {
                MessageBox.Show("用户名或密码错误!", "登录失败");
            }
        }
    }
}
