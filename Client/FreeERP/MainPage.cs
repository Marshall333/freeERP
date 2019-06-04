using FreeERP.View;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace FreeERP
{
    public partial class MainPage : Form
    {
        
        public MainPage()
        {
            InitializeComponent();
            this.Visible = false;

            //loginForm = new Login(this);
            //loginForm.ShowDialog();
        }

        private void repoManagerForm_Load(object sender, EventArgs e)
        {

        }
    }
}
